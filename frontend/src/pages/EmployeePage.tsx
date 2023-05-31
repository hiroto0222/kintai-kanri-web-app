import { Container, Stack, Typography } from "@mui/material";
import { Helmet } from "react-helmet-async";
import { useTranslation } from "react-i18next";
import { useParams } from "react-router-dom";
import useGetEmployee from "../hooks/api/useGetEmployee";
import usePrivateAxios from "../hooks/usePrivateAxios";

const EmployeePage = () => {
  const { employeeId } = useParams();
  const { t } = useTranslation();
  const privateAxios = usePrivateAxios();
  const { employee } = useGetEmployee(privateAxios, employeeId);

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
            {t("nav.employees")} | {employee?.last_name} {employee?.first_name}
          </Typography>
        </Stack>
      </Container>
    </>
  );
};

export default EmployeePage;
