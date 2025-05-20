import axios from 'axios';
import type { Product, OrderRequest, Order } from '../types/api';

const API_URL = 'http://localhost:8080/api';
const API_KEY = 'apitest';

const api = axios.create({
    baseURL: API_URL,
    headers: {
        'Content-Type': 'application/json',
    },
});

export const getProducts = async (): Promise<Product[]> => {
    const response = await api.get('/product');
    return response.data;
};

export const placeOrder = async (order: OrderRequest): Promise<Order> => {
    const response = await api.post('/order', order, {
        headers: {
            'api_key': API_KEY,
        },
    });
    return response.data;
};
