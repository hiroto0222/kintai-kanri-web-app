import {
  Button,
  Card,
  CardContent,
  Container,
  Typography,
} from "@mui/material";
import { Helmet } from "react-helmet-async";
import { SubmitHandler, useForm } from "react-hook-form";
import LoginForm, { LoginFormProps } from "../components/auth/LoginForm";
import useAuthApi from "../hooks/api/useAuthApi";

const LoginPage = () => {
  const { login } = useAuthApi();
  const { handleSubmit, control } = useForm<LoginFormProps>({
    mode: "onBlur",
    criteriaMode: "all",
    shouldFocusError: true,
  });

  const onSubmit: SubmitHandler<LoginFormProps> = (data) => {
    login(data.Email, data.Password);
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
          component="form"
          onSubmit={handleSubmit(onSubmit)}
        >
          <CardContent>
            <Typography variant="h4" align="center" gutterBottom>
              Welcome
            </Typography>
            <LoginForm control={control} />
            <Button fullWidth size="large" type="submit" variant="contained">
              Login
            </Button>
          </CardContent>
        </Card>
      </Container>
    </>
  );
};

export default LoginPage;
