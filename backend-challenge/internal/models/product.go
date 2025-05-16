package models

// Product represents a product available for order
type Product struct {
	ID       string       `json:"id"`
	Name     string       `json:"name"`
	Price    float64      `json:"price"`
	Category string       `json:"category"`
	Image    ProductImage `json:"image"`
}

// ProductImage represents various image sizes for a product
type ProductImage struct {
	Thumbnail string `json:"thumbnail"`
	Mobile    string `json:"mobile"`
	Tablet    string `json:"tablet"`
	Desktop   string `json:"desktop"`
}

// MockProducts returns a list of sample products for demo purposes
func MockProducts() []Product {
	return []Product{
		{
			ID:       "1",
			Name:     "Chicken Waffle",
			Price:    13.3,
			Category: "Waffle",
			Image: ProductImage{
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-waffle-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-waffle-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-waffle-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-waffle-desktop.jpg",
			},
		},
		{
			ID:       "2",
			Name:     "Beef Burger",
			Price:    15.5,
			Category: "Burger",
			Image: ProductImage{
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-burger-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-burger-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-burger-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-burger-desktop.jpg",
			},
		},
		{
			ID:       "3",
			Name:     "Margherita Pizza",
			Price:    12.0,
			Category: "Pizza",
			Image: ProductImage{
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-pizza-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-pizza-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-pizza-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-pizza-desktop.jpg",
			},
		},
	}
}
