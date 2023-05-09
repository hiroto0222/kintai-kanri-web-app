import {
  Button,
  Card,
  CardContent,
  Container,
  Stack,
  TextField,
  Typography,
} from "@mui/material";
import { useState } from "react";
import { Helmet } from "react-helmet-async";
import useAuthApi from "../hooks/api/useAuthApi";

const LoginPage = () => {
  const { login } = useAuthApi();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    e.preventDefault();
    login(email, password);
  };

  return (
    <>
      <Helmet>
        <title>Login</title>
      </Helmet>
      <Container
        component="main"
        maxWidth="sm"
        sx={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          minHeight: "100vh",
        }}
      >
        <Card
          sx={{
            width: "100%",
          }}
        >
          <CardContent>
            <Typography variant="h4" align="center" gutterBottom>
              Welcome
            </Typography>
            <Stack spacing={3} padding={3}>
              <TextField
                name="email"
                label="Email address"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
              <TextField
                name="password"
                label="Password"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
              <Button
                fullWidth
                size="large"
                type="submit"
                variant="contained"
                onClick={(e) => handleSubmit(e)}
              >
                Login
              </Button>
            </Stack>
          </CardContent>
        </Card>
      </Container>
    </>
  );
};

export default LoginPage;