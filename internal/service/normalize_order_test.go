package service

import (
	"technical-skills/internal/model"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNormalizeOrder(t *testing.T) {
	tests := []struct {
		name     string
		input    []model.InputOrder
		expected []model.CleanedOrder
	}{
		{
			name: "Case 1",
			input: []model.InputOrder{
				{
					No:                1,
					PlatformProductId: "FG0A-CLEAR-IPHONE16PROMAX",
					Qty:               2,
					UnitPrice:         50,
					TotalPrice:        100,
				},
			},
			expected: []model.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "IPHONE16PROMAX",
					Qty:        2,
					UnitPrice:  50,
					TotalPrice: 100,
				},
				{
					No:         2,
					ProductId:  "WIPING-CLOTH",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         3,
					ProductId:  "CLEAR-CLEANNER",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
			},
		},
		{
			name: "Case 2",
			input: []model.InputOrder{
				{
					No:                1,
					PlatformProductId: "x2-3&FG0A-CLEAR-IPHONE16PROMAX",
					Qty:               2,
					UnitPrice:         50,
					TotalPrice:        100,
				},
			},
			expected: []model.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "IPHONE16PROMAX",
					Qty:        2,
					UnitPrice:  50,
					TotalPrice: 100,
				},
				{
					No:         2,
					ProductId:  "WIPING-CLOTH",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         3,
					ProductId:  "CLEAR-CLEANNER",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
			},
		},
		{
			name: "Case 3",
			input: []model.InputOrder{
				{
					No:                1,
					PlatformProductId: "x2-3&FG0A-CLEAR-IPHONE16PROMAX*3",
					Qty:               1,
					UnitPrice:         90,
					TotalPrice:        90,
				},
			},
			expected: []model.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "IPHONE16PROMAX",
					Qty:        3,
					UnitPrice:  30,
					TotalPrice: 90,
				},
				{
					No:         2,
					ProductId:  "WIPING-CLOTH",
					Qty:        3,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         3,
					ProductId:  "CLEAR-CLEANNER",
					Qty:        3,
					UnitPrice:  0,
					TotalPrice: 0,
				},
			},
		},
		{
			name: "Case 4",
			input: []model.InputOrder{
				{
					No:                1,
					PlatformProductId: "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B",
					Qty:               1,
					UnitPrice:         80,
					TotalPrice:        80,
				},
			},
			expected: []model.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-OPPOA3",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
				{
					No:         2,
					ProductId:  "FG0A-CLEAR-OPPOA3-B",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3-B",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
				{
					No:         3,
					ProductId:  "WIPING-CLOTH",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         4,
					ProductId:  "CLEAR-CLEANNER",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
			},
		},
		{
			name: "Case 5",
			input: []model.InputOrder{
				{
					No:                1,
					PlatformProductId: "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B/FG0A-MATTE-OPPOA3",
					Qty:               1,
					UnitPrice:         120,
					TotalPrice:        120,
				},
			},
			expected: []model.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-OPPOA3",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
				{
					No:         2,
					ProductId:  "FG0A-CLEAR-OPPOA3-B",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3-B",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
				{
					No:         3,
					ProductId:  "FG0A-MATTE-OPPOA3",
					MaterialId: "FG0A-MATTE",
					ModelId:    "OPPOA3",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
				{
					No:         4,
					ProductId:  "WIPING-CLOTH",
					Qty:        3,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         5,
					ProductId:  "CLEAR-CLEANNER",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         6,
					ProductId:  "MATTE-CLEANNER",
					Qty:        1,
					UnitPrice:  0,
					TotalPrice: 0,
				},
			},
		},
		{
			name: "Case 6",
			input: []model.InputOrder{
				{
					No:                1,
					PlatformProductId: "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3",
					Qty:               1,
					UnitPrice:         120,
					TotalPrice:        120,
				},
			},
			expected: []model.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-OPPOA3",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3",
					Qty:        2,
					UnitPrice:  40,
					TotalPrice: 80,
				},
				{
					No:         2,
					ProductId:  "FG0A-MATTE-OPPOA3",
					MaterialId: "FG0A-MATTE",
					ModelId:    "OPPOA3",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
				{
					No:         3,
					ProductId:  "WIPING-CLOTH",
					Qty:        3,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         4,
					ProductId:  "CLEAR-CLEANNER",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         5,
					ProductId:  "MATTE-CLEANNER",
					Qty:        1,
					UnitPrice:  0,
					TotalPrice: 0,
				},
			},
		},
		{
			name: "Case 7",
			input: []model.InputOrder{
				{
					No:                1,
					PlatformProductId: "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3*2",
					Qty:               1,
					UnitPrice:         160,
					TotalPrice:        160,
				},
				{
					No:                2,
					PlatformProductId: "FG0A-PRIVACY-IPHONE16PROMAX",
					Qty:               1,
					UnitPrice:         50,
					TotalPrice:        50,
				},
			},
			expected: []model.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-OPPOA3",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3",
					Qty:        2,
					UnitPrice:  40,
					TotalPrice: 80,
				},
				{
					No:         2,
					ProductId:  "FG0A-MATTE-OPPOA3",
					MaterialId: "FG0A-MATTE",
					ModelId:    "OPPOA3",
					Qty:        2,
					UnitPrice:  40,
					TotalPrice: 80,
				},
				{
					No:         3,
					ProductId:  "FG0A-PRIVACY-IPHONE16PROMAX",
					MaterialId: "FG0A-PRIVACY",
					ModelId:    "IPHONE16PROMAX",
					Qty:        1,
					UnitPrice:  50,
					TotalPrice: 50,
				},
				{
					No:         4,
					ProductId:  "WIPING-CLOTH",
					Qty:        5,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         5,
					ProductId:  "CLEAR-CLEANNER",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         6,
					ProductId:  "MATTE-CLEANNER",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         7,
					ProductId:  "PRIVACY-CLEANNER",
					Qty:        1,
					UnitPrice:  0,
					TotalPrice: 0,
				},
			},
		},
		{
			name: "Case 8 (Extra)",
			input: []model.InputOrder{
				{
					No:                1,
					PlatformProductId: "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3*2---",
					Qty:               1,
					UnitPrice:         160,
					TotalPrice:        160,
				},
				{
					No:                2,
					PlatformProductId: "  ---FG0A-PRIVACY-IPHONE16PROMAX*3---  ",
					Qty:               1,
					UnitPrice:         150,
					TotalPrice:        150,
				},
				{
					No:                3,
					PlatformProductId: "FG0A-PRIVACY-SAMSUNGS25*3",
					Qty:               1,
					UnitPrice:         120,
					TotalPrice:        120,
				},
			},
			expected: []model.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-OPPOA3",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3",
					Qty:        2,
					UnitPrice:  40,
					TotalPrice: 80,
				},
				{
					No:         2,
					ProductId:  "FG0A-MATTE-OPPOA3",
					MaterialId: "FG0A-MATTE",
					ModelId:    "OPPOA3",
					Qty:        2,
					UnitPrice:  40,
					TotalPrice: 80,
				},
				{
					No:         3,
					ProductId:  "FG0A-PRIVACY-IPHONE16PROMAX",
					MaterialId: "FG0A-PRIVACY",
					ModelId:    "IPHONE16PROMAX",
					Qty:        3,
					UnitPrice:  50,
					TotalPrice: 150,
				},
				{
					No:         4,
					ProductId:  "FG0A-PRIVACY-SAMSUNGS25",
					MaterialId: "FG0A-PRIVACY",
					ModelId:    "SAMSUNGS25",
					Qty:        3,
					UnitPrice:  40,
					TotalPrice: 120,
				},
				{
					No:         5,
					ProductId:  "WIPING-CLOTH",
					Qty:        10,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         6,
					ProductId:  "CLEAR-CLEANNER",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         7,
					ProductId:  "MATTE-CLEANNER",
					Qty:        2,
					UnitPrice:  0,
					TotalPrice: 0,
				},
				{
					No:         8,
					ProductId:  "PRIVACY-CLEANNER",
					Qty:        6,
					UnitPrice:  0,
					TotalPrice: 0,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeOrder(tt.input)
			if diff := cmp.Diff(tt.expected, result); diff != "" {
				t.Errorf("Mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
