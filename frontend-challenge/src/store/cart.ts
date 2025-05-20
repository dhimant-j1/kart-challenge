import { create } from 'zustand';
import type { CartItem, Product } from '../types/api';

interface CartStore {
    items: CartItem[];
    addItem: (product: Product) => void;
    removeItem: (productId: string) => void;
    updateQuantity: (productId: string, quantity: number) => void;
    clearCart: () => void;
    total: number;
}

export const useCartStore = create<CartStore>((set) => ({
    items: [],
    total: 0,
    
    addItem: (product: Product) => {
        set((state) => {
            const existingItem = state.items.find(item => item.productId === product.id);
            
            if (existingItem) {
                return {
                    ...state,
                    items: state.items.map(item =>
                        item.productId === product.id
                            ? { ...item, quantity: item.quantity + 1 }
                            : item
                    ),
                    total: state.total + product.price,
                };
            }
            
            return {
                ...state,
                items: [...state.items, { productId: product.id, quantity: 1, product }],
                total: state.total + product.price,
            };
        });
    },
    
    removeItem: (productId: string) => {
        set((state) => {
            const item = state.items.find(item => item.productId === productId);
            if (!item) return state;
            
            return {
                ...state,
                items: state.items.filter(item => item.productId !== productId),
                total: state.total - (item.product.price * item.quantity),
            };
        });
    },
    
    updateQuantity: (productId: string, quantity: number) => {
        set((state) => {
            const item = state.items.find(item => item.productId === productId);
            if (!item) return state;
            
            const quantityDiff = quantity - item.quantity;
            
            return {
                ...state,
                items: state.items.map(item =>
                    item.productId === productId
                        ? { ...item, quantity }
                        : item
                ),
                total: state.total + (item.product.price * quantityDiff),
            };
        });
    },
    
    clearCart: () => {
        set({ items: [], total: 0 });
    },
}));
