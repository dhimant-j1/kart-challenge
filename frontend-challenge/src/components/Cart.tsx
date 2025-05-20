import React, { useState } from "react";
import {
  Box,
  Paper,
  Typography,
  List,
  ListItem,
  Button,
  IconButton,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Divider,
} from "@mui/material";
import { Close } from "@mui/icons-material"; // Updated to use Close icon
import { useCartStore } from "../store/cart";
import { placeOrder } from "../services/api";
import { toast } from "react-hot-toast";

export const Cart: React.FC = () => {
  const { items, total, removeItem, clearCart } = useCartStore();
  const [isOrderDialogOpen, setIsOrderDialogOpen] = useState(false);
  const [isProcessing, setIsProcessing] = useState(false);

  const handlePlaceOrder = async () => {
    if (items.length === 0) return;

    setIsProcessing(true);
    try {
      await placeOrder({
        items: items.map(({ productId, quantity }) => ({
          productId,
          quantity,
        })),
      });

      setIsOrderDialogOpen(true);
      clearCart();
    } catch (error) {
      toast.error("Failed to place order. Please try again.");
    }
    setIsProcessing(false);
  };

  return (
    <>
      <Paper
        sx={{
          p: 3,
          height: "100%",
          display: "flex",
          flexDirection: "column",
          borderRadius: 2,
          boxShadow: "0px 4px 12px rgba(0, 0, 0, 0.05)",
        }}
      >
        <Typography
          variant="h6"
          gutterBottom
          sx={{
            fontWeight: 600,
            fontSize: "1.25rem",
            color: "#1F2937",
            mb: 3,
          }}
        >
          Shopping Cart
        </Typography>

        <List
  sx={{
    flexGrow: 1,
    overflowY: "auto",
    pr: 1, // Add padding to avoid content being cut off
    scrollbarWidth: "none", // For Firefox
    "&::-webkit-scrollbar": {
      display: "none", // For WebKit browsers (Chrome, Safari)
    },
  }}
>
  {items.length === 0 ? (
    <Typography color="text.secondary" align="center" sx={{ mt: 4 }}>
      Your cart is empty
    </Typography>
  ) : (
    items.map((item) => (
      <React.Fragment key={item.productId}>
        <ListItem
          sx={{
            py: 2,
            px: 0,
            display: "flex",
            alignItems: "center", // Center-align content vertically
            justifyContent: "space-between", // Ensure proper spacing between elements
            gap: 2,
          }}
        >
          <Box sx={{ flex: 1 }}>
            <Typography
              variant="subtitle1"
              sx={{
                fontWeight: 500,
                color: "#1F2937",
                mb: 0.5,
                textAlign: "left", // Left-align the product name
              }}
            >
              {item.product.name}
            </Typography>
            <Typography variant="body2" sx={{ color: "#6B7280" }}>
              ${item.product.price.toFixed(2)} × {item.quantity}
            </Typography>
          </Box>
          <Typography
            variant="subtitle1"
            sx={{
              fontWeight: 600,
              color: "#1F2937",
              mr: 2, // Add margin to separate the price from the close button
            }}
          >
            ${(item.quantity * item.product.price).toFixed(2)}
          </Typography>
          <IconButton
            edge="end"
            aria-label="delete"
            onClick={() => removeItem(item.productId)}
            sx={{
              color: "#9CA3AF",
              "&:hover": {
                color: "#EF4444",
              },
            }}
          >
            <Close /> {/* Updated to Close icon */}
          </IconButton>
        </ListItem>
        <Divider sx={{ my: 1 }} />
      </React.Fragment>
    ))
  )}
</List>

        <Box sx={{ mt: 3 }}>
          <Box
            sx={{
              display: "flex",
              justifyContent: "space-between",
              mb: 1,
            }}
          >
            <Typography sx={{ color: "#6B7280" }}>Subtotal:</Typography>
            <Typography sx={{ fontWeight: 500 }}>
              ${total.toFixed(2)}
            </Typography>
          </Box>
          <Box
            sx={{
              display: "flex",
              justifyContent: "space-between",
              mb: 2,
            }}
          >
            <Typography sx={{ color: "#6B7280" }}>
              Carbon-neutral delivery:
            </Typography>
            <Typography sx={{ fontWeight: 500 }}>$0.00</Typography>
          </Box>
          <Box
            sx={{
              display: "flex",
              justifyContent: "space-between",
              mb: 3,
            }}
          >
            <Typography sx={{ fontWeight: 600 }}>Total:</Typography>
            <Typography
              sx={{
                fontWeight: 600,
                color: "#EF4444",
              }}
            >
              ${total.toFixed(2)}
            </Typography>
          </Box>
          <Button
            variant="contained"
            fullWidth
            disabled={items.length === 0 || isProcessing}
            onClick={handlePlaceOrder}
            sx={{
              textTransform: "none",
              py: 1.5,
              bgcolor: "#EF4444",
              "&:hover": {
                bgcolor: "#DC2626",
              },
              "&:disabled": {
                bgcolor: "#F3F4F6",
              },
            }}
          >
            {isProcessing ? "Processing..." : "Confirm Order"}
          </Button>
        </Box>
      </Paper>

      <Dialog
        open={isOrderDialogOpen}
        onClose={() => setIsOrderDialogOpen(false)}
        PaperProps={{
          sx: {
            borderRadius: 2,
            p: 2,
          },
        }}
      >
        <DialogTitle sx={{ fontWeight: 600, color: "#1F2937" }}>
          Order Confirmed!
        </DialogTitle>
        <DialogContent>
          <Typography>
            Thank you for your order. It will be delivered soon!
          </Typography>
        </DialogContent>
        <DialogActions>
          <Button
            onClick={() => setIsOrderDialogOpen(false)}
            sx={{
              color: "#EF4444",
              "&:hover": {
                bgcolor: "rgba(239, 68, 68, 0.04)",
              },
            }}
          >
            Close
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};