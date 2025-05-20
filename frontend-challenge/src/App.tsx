import Container from "@mui/material/Container";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import CircularProgress from "@mui/material/CircularProgress";
import CssBaseline from "@mui/material/CssBaseline";
import Typography from "@mui/material/Typography";
import { Toaster } from "react-hot-toast";
import { useQuery } from "@tanstack/react-query";
import { ProductCard } from "./components/ProductCard";
import { Cart } from "./components/Cart";
import { getProducts } from "./services/api";
import "./App.css";

function App() {
  const {
    data: products,
    isLoading,
    error,
  } = useQuery({
    queryKey: ["products"],
    queryFn: getProducts,
  });

  if (isLoading) {
    return (
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        minHeight="100vh"
      >
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return (
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        minHeight="100vh"
      >
        <Typography color="error">
          Error loading products. Please try again later.
        </Typography>
      </Box>
    );
  }

  return (
    <>
      <CssBaseline />
      <Container
        maxWidth={false}
        sx={{ px: { xs: 2, sm: 4, md: 6, lg: 8 }, pt: 4, pb: 8 }}
      >
        <Typography
          variant="h4"
          component="h1"
          sx={{ mb: 4, fontWeight: "bold", color: "#2D2D2D" }}
        >
          Desserts
        </Typography>
        <Grid container spacing={3}>
        <Grid item xs={12} lg={8} xl={9}>
          <Grid container spacing={2}>
            {products?.map((product) => (
              <Grid item xs={12} sm={6} md={4} lg={3} key={product.id}>
                <ProductCard product={product} />
              </Grid>
            ))}
          </Grid>
        </Grid>
        <Grid item xs={12} lg={4} xl={3} sx={{ position: "sticky", top: 20 }}>
          <Box
            sx={{
              bgcolor: "white",
              p: 3,
              borderRadius: 2,
              boxShadow: "0px 4px 12px rgba(0, 0, 0, 0.05)",
            }}
          >
            <Cart />
          </Box>
        </Grid>
      </Grid>
      </Container>
      <Toaster position="bottom-right" />
    </>
  );
}

export default App;
