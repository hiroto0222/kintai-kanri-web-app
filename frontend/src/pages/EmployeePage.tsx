import { Container, Grid, Stack, Typography } from "@mui/material";
import { Helmet } from "react-helmet-async";
import { useTranslation } from "react-i18next";
import { useParams } from "react-router-dom";
import ClockInsList from "../components/dashboard/ClockInsList";
import Loading from "../components/dashboard/Loading";
import useGetEmployee from "../hooks/api/useGetEmployee";
import useListClockIns from "../hooks/api/useListClockIns";
import usePrivateAxios from "../hooks/usePrivateAxios";

const EmployeePage = () => {
  const { employeeId } = useParams();
  const { t } = useTranslation();
  const privateAxios = usePrivateAxios();
  const { employee } = useGetEmployee(privateAxios, employeeId);
  const { clockIns, loading } = useListClockIns(privateAxios, employeeId);

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
        {loading ? (
          <Loading />
        ) : (
          <Grid item xs={12}>
            <ClockInsList data={clockIns} />
          </Grid>
        )}
      </Container>
    </>
  );
};

export default EmployeePage;
