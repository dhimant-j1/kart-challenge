import React from "react";
import {
  Card,
  CardMedia,
  CardContent,
  Typography,
  Button,
  Box,
  IconButton,
} from "@mui/material";
import { Add, Remove } from "@mui/icons-material";
import type { Product } from "../types/api";
import { useCartStore } from "../store/cart";

interface ProductCardProps {
  product: Product;
}

export const ProductCard: React.FC<ProductCardProps> = ({ product }) => {
  const { items, addItem, updateQuantity } = useCartStore();
  const cartItem = items.find((item) => item.productId === product.id);

  return (
    <Card
      elevation={0}
      sx={{
        display: "flex",
        flexDirection: "column",
        height: "100%",
        borderRadius: 2,
        border: "1px solid #E5E7EB",
        borderBottom: "none", // Remove bottom border
        position: "relative",
        overflow: "visible",
        transition: "transform 0.2s, box-shadow 0.2s",
        "&:hover": {
          transform: "translateY(-4px)",
          boxShadow: "0px 8px 16px rgba(0, 0, 0, 0.1)",
        },
      }}
    >
      <CardMedia
        component="img"
        image={product.image.desktop}
        alt={product.name}
        sx={{
          height: 240,
          borderRadius: "8px 8px 0 0",
          objectFit: "cover",
          bgcolor: "#F3F4F6", // Light gray background for placeholder
        }}
        onError={(e) => {
          const target = e.target as HTMLImageElement;
          target.src = "https://placehold.co/502x480/F3F4F6/6B7280?text=No+Image"; // Fallback image
        }}
      />
      {/* Button positioned directly below the image */}
      <Box
        sx={{
          position: "absolute",
          top: 220, // Position the button directly below the image
          left: "50%",
          transform: "translateX(-50%)",
          display: "flex",
          justifyContent: "center",
          zIndex: 1,
        }}
      >
        {!cartItem ? (
          <Button
            variant="contained"
            onClick={() => addItem(product)}
            sx={{
              borderRadius: "50px", // Pill shape
              textTransform: "none",
              py: 1,
              px: 3,
              bgcolor: "#EF4444", // Red background
              color: "white",
              boxShadow: "0px 4px 12px rgba(0, 0, 0, 0.1)",
              "&:hover": {
                bgcolor: "#DC2626", // Darker red on hover
              },
            }}
          >
            Add to Cart
          </Button>
        ) : (
          <Box
            sx={{
              display: "flex",
              alignItems: "center",
              gap: 1,
              bgcolor: "#EF4444", // Red background for quantity controls
              borderRadius: "50px", // Pill shape
              p: 0.5,
              py: 1,
              px: 2,
            }}
          >
            <IconButton
              size="small"
              onClick={() =>
                cartItem.quantity > 1
                  ? updateQuantity(product.id, cartItem.quantity - 1)
                  : updateQuantity(product.id, 0) // Set to 0 if quantity is 1
              }
              sx={{
                color: "white", // White icon color
                "&:hover": {
                  bgcolor: "#DC2626", // Darker red on hover
                },
              }}
            >
              <Remove fontSize="small" />
            </IconButton>
            <Typography
              sx={{
                flex: 1,
                textAlign: "center",
                fontWeight: 500,
                color: "white", // White text color
              }}
            >
              {cartItem.quantity}
            </Typography>
            <IconButton
              size="small"
              onClick={() => updateQuantity(product.id, cartItem.quantity + 1)}
              sx={{
                color: "white", // White icon color
                "&:hover": {
                  bgcolor: "#DC2626", // Darker red on hover
                },
              }}
            >
              <Add fontSize="small" />
            </IconButton>
          </Box>
        )}
      </Box>
      <CardContent sx={{ p: 3, flex: 1 }}>
        <Typography
          variant="subtitle1"
          sx={{
            fontSize: "0.875rem",
            color: "#6B7280",
            mb: 1,
            fontWeight: 500,
          }}
        >
          {product.category}
        </Typography>
        <Typography
          variant="h6"
          sx={{
            fontSize: "1.125rem",
            fontWeight: 600,
            color: "#1F2937",
            mb: 2,
            minHeight: "3.5rem",
            display: "-webkit-box",
            WebkitLineClamp: 2,
            WebkitBoxOrient: "vertical",
            overflow: "hidden",
          }}
        >
          {product.name}
        </Typography>
        <Typography
          variant="body1"
          sx={{
            color: "#EF4444",
            fontWeight: 600,
            fontSize: "1.125rem",
            mb: 2,
          }}
        >
          ${product.price.toFixed(2)}
        </Typography>
      </CardContent>
    </Card>
  );
};