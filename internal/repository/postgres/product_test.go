package postgres_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/satriowisnugroho/catalog/internal/entity"
	"github.com/satriowisnugroho/catalog/internal/entity/types"
	"github.com/satriowisnugroho/catalog/internal/repository/postgres"
	"github.com/satriowisnugroho/catalog/test/fixture"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	testcases := []struct {
		name      string
		ctx       context.Context
		input     *entity.Product
		createErr error
		wantErr   bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:      "fail exec query",
			ctx:       context.Background(),
			input:     &entity.Product{},
			createErr: errors.New("fail exec"),
			wantErr:   true,
		},
		{
			name:    "success",
			ctx:     context.Background(),
			input:   &entity.Product{},
			wantErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			if tc.createErr != nil {
				mock.ExpectQuery("^INSERT INTO(.+)").WillReturnError(tc.createErr)
			} else {
				row := sqlmock.NewRows([]string{"id"})
				result := row.AddRow(1)
				mock.ExpectQuery("^INSERT INTO(.+)").WillReturnRows(result)
			}

			dbx := sqlx.NewDb(db, "mock")
			repo := postgres.NewProductRepository(dbx)

			err = repo.CreateProduct(tc.ctx, tc.input)
			assert.Equal(t, tc.wantErr, err != nil)
			if !tc.wantErr {
				assert.Equal(t, 1, tc.input.ID)
			}
		})
	}
}

func TestGetProductByID(t *testing.T) {
	testcases := []struct {
		name      string
		ctx       context.Context
		fetchErr  error
		fetchRows []string
		expected  *entity.Product
		wantErr   bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:     "fail fetch query error",
			ctx:      context.Background(),
			fetchErr: errors.New("fail fetch"),
			wantErr:  true,
		},
		{
			name:      "fail fetch return error rows",
			ctx:       context.Background(),
			fetchRows: []string{"unknown_column"},
			wantErr:   true,
		},
		{
			name:      "record not found",
			ctx:       context.Background(),
			fetchRows: postgres.ProductColumns,
			wantErr:   true,
		},
		{
			name:      "success",
			ctx:       context.Background(),
			fetchRows: postgres.ProductColumns,
			expected:  &entity.Product{Tenant: types.TenantLoremType},
			wantErr:   false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			if tc.fetchErr != nil {
				mock.ExpectQuery("^SELECT(.+)").WillReturnError(tc.fetchErr)
			} else {
				rows := sqlmock.NewRows(tc.fetchRows)
				if tc.expected != nil {
					rows = rows.AddRow(
						tc.expected.ID,
						tc.expected.SKU,
						tc.expected.Title,
						tc.expected.Category,
						tc.expected.Condition,
						tc.expected.Tenant,
						tc.expected.Qty,
						tc.expected.Price,
						tc.expected.CreatedAt,
						tc.expected.UpdatedAt,
					)
				} else if len(tc.fetchRows) == 1 {
					rows = rows.AddRow(1)
				}

				mock.ExpectQuery("^SELECT(.+)").WillReturnRows(rows)
			}

			dbx := sqlx.NewDb(db, "mock")
			repo := postgres.NewProductRepository(dbx)
			result, err := repo.GetProductByID(tc.ctx, 123)
			assert.Equal(t, tc.wantErr, err != nil, err)
			if !tc.wantErr {
				assert.EqualValues(t, tc.expected, result)
			}
		})
	}
}

func TestGetProducts(t *testing.T) {
	testcases := []struct {
		name      string
		ctx       context.Context
		payload   *entity.GetProductPayload
		fetchErr  error
		fetchRows []string
		expected  []*entity.Product
		wantErr   bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:     "fail fetch query error",
			ctx:      context.Background(),
			fetchErr: errors.New("fail fetch"),
			payload:  &entity.GetProductPayload{},
			wantErr:  true,
		},
		{
			name:      "fail fetch return error rows",
			ctx:       context.Background(),
			payload:   &entity.GetProductPayload{},
			fetchRows: []string{"unknown_column"},
			wantErr:   true,
		},
		{
			name:      "success when limit greater than 100",
			ctx:       context.Background(),
			payload:   &entity.GetProductPayload{Limit: 99999},
			fetchRows: postgres.ProductColumns,
			expected:  []*entity.Product{{Tenant: types.TenantLoremType}},
			wantErr:   false,
		},
		{
			name: "success",
			ctx:  context.Background(),
			payload: &entity.GetProductPayload{
				SKU:          "SKU123",
				TitleKeyword: "Product",
				Category:     "foo",
				Condition:    1,
				Tenant:       types.TenantLoremType,
				OrderBy:      "created_at DESC",
				Offset:       0,
				Limit:        10,
			},
			fetchRows: postgres.ProductColumns,
			expected:  []*entity.Product{{Tenant: types.TenantLoremType}},
			wantErr:   false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			if tc.fetchErr != nil {
				mock.ExpectQuery("^SELECT(.+)").WillReturnError(tc.fetchErr)
			} else {
				rows := sqlmock.NewRows(tc.fetchRows)
				if tc.expected != nil {
					rows = rows.AddRow(
						tc.expected[0].ID,
						tc.expected[0].SKU,
						tc.expected[0].Title,
						tc.expected[0].Category,
						tc.expected[0].Condition,
						tc.expected[0].Tenant,
						tc.expected[0].Qty,
						tc.expected[0].Price,
						tc.expected[0].CreatedAt,
						tc.expected[0].UpdatedAt,
					)
				} else if len(tc.fetchRows) == 1 {
					rows = rows.AddRow(1)
				}

				mock.ExpectQuery("^SELECT(.+)").WillReturnRows(rows)
			}

			dbx := sqlx.NewDb(db, "mock")
			repo := postgres.NewProductRepository(dbx)
			result, err := repo.GetProducts(tc.ctx, tc.payload)
			assert.Equal(t, tc.wantErr, err != nil, err)
			if !tc.wantErr {
				assert.EqualValues(t, tc.expected, result)
			}
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	testcases := []struct {
		name      string
		ctx       context.Context
		updateErr error
		wantErr   bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:      "fail exec query",
			ctx:       context.Background(),
			updateErr: errors.New("fail exec"),
			wantErr:   true,
		},
		{
			name:    "success",
			ctx:     context.Background(),
			wantErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			if tc.updateErr != nil {
				mock.ExpectExec("^UPDATE products(.+)").WillReturnError(tc.updateErr)
			} else {
				mock.ExpectExec("^UPDATE products(.+)").WillReturnResult(sqlmock.NewResult(1, 1))
			}

			dbx := sqlx.NewDb(db, "mock")
			repo := postgres.NewProductRepository(dbx)
			err = repo.UpdateProduct(tc.ctx, &entity.Product{})
			assert.Equal(t, tc.wantErr, err != nil, err)
		})
	}
}
