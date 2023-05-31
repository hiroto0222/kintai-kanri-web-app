import {
  Button,
  Card,
  CardContent,
  Container,
  Typography,
} from "@mui/material";
import { Helmet } from "react-helmet-async";
import { SubmitHandler, useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";
import LoginForm, { LoginFormProps } from "../components/auth/LoginForm";
import useAuthApi from "../hooks/api/useAuthApi";
import usePrivateAxios from "../hooks/usePrivateAxios";

const LoginPage = () => {
  const { t } = useTranslation();
  const privateAxios = usePrivateAxios();
  const { login } = useAuthApi(privateAxios);
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
              {t("auth.welcome")}
            </Typography>
            <LoginForm control={control} />
            <Button fullWidth size="large" type="submit" variant="contained">
              Login
            </Button>
            <Typography paddingTop={4} variant="body1">
              テスト管理者アカウント:
              <br />
              email: admin@email.com <br />
              password: admin123
            </Typography>
          </CardContent>
        </Card>
      </Container>
    </>
  );
};

export default LoginPage;
