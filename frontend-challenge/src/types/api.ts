export interface Product {
    id: string;
    name: string;
    price: number;
    category: string;
    image: {
        thumbnail: string;
        mobile: string;
        tablet: string;
        desktop: string;
    };
}

export interface OrderItem {
    productId: string;
    quantity: number;
}

export interface OrderRequest {
    items: OrderItem[];
    couponCode?: string;
}

export interface Order {
    id: string;
    total: number;
    discounts: number;
    items: OrderItem[];
    products: Product[];
}

export interface CartItem extends OrderItem {
    product: Product;
}
