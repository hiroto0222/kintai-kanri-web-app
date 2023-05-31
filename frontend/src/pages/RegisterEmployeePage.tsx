import {
  Button,
  Card,
  CardContent,
  Container,
  Grid,
  Stack,
  Typography,
} from "@mui/material";
import { Helmet } from "react-helmet-async";
import { SubmitHandler, useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";
import RegisterForm, {
  RegisterFormProps,
} from "../components/auth/RegisterForm";
import useAuthApi from "../hooks/api/useAuthApi";
import usePrivateAxios from "../hooks/usePrivateAxios";

const RegisterPage = () => {
  const { t } = useTranslation();
  const privateAxios = usePrivateAxios();
  const { register } = useAuthApi(privateAxios);
  const navigate = useNavigate();

  const { handleSubmit, control } = useForm<RegisterFormProps>({
    mode: "onBlur",
    criteriaMode: "all",
    shouldFocusError: true,
  });

  const onSubmit: SubmitHandler<RegisterFormProps> = (data) => {
    register(data).then(() => {
      navigate("/dashboard/employees");
    });
  };

  return (
    <>
      <Helmet>
        <title>{t("nav.employees")}</title>
      </Helmet>
      <Container maxWidth="xl">
        <Stack
          direction="row"
          alignItems="center"
          justifyContent="space-between"
          sx={{ mb: 5 }}
        >
          <Typography variant="h3">
            {t("dashboard.register_employee")}
          </Typography>
        </Stack>
        <Grid justifyContent="space-evenly" container spacing={3}>
          <Grid item xs={12} md={10} lg={8}>
            <Card component="form" onSubmit={handleSubmit(onSubmit)}>
              <CardContent>
                <RegisterForm control={control} />
                <Button
                  fullWidth
                  size="large"
                  type="submit"
                  variant="contained"
                >
                  Register
                </Button>
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      </Container>
    </>
  );
};

export default RegisterPage;
