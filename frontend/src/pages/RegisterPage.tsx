import {
  Button,
  Card,
  CardContent,
  Container,
  Typography,
} from "@mui/material";
import { Helmet } from "react-helmet-async";
import { SubmitHandler, useForm } from "react-hook-form";
import RegisterForm, {
  RegisterFormProps,
} from "../components/auth/RegisterForm";
import useAuthApi from "../hooks/api/useAuthApi";

const RegisterPage = () => {
  const { register } = useAuthApi();
  const { handleSubmit, control } = useForm<RegisterFormProps>({
    mode: "onBlur",
    criteriaMode: "all",
    shouldFocusError: true,
  });

  const onSubmit: SubmitHandler<RegisterFormProps> = (data) => {
    register(data);
  };

  return (
    <>
      <Helmet>
        <title>Register</title>
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
              Register
            </Typography>
            <RegisterForm control={control} />
            <Button fullWidth size="large" type="submit" variant="contained">
              Sign Up
            </Button>
          </CardContent>
        </Card>
      </Container>
    </>
  );
};

export default RegisterPage;
