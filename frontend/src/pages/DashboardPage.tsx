import { Container, Grid, Stack, Typography, useTheme } from "@mui/material";
import { useContext } from "react";
import { Helmet } from "react-helmet-async";
import { useTranslation } from "react-i18next";
import CustomButton from "../components/dashboard/Button";
import Clock from "../components/dashboard/Clock";
import ClockInsList from "../components/dashboard/ClockInsList";
import Loading from "../components/dashboard/Loading";
import { authContext } from "../context/auth";
import useClockInsClockOutsApi from "../hooks/api/useClockInsClockOutsApi";
import useListClockIns from "../hooks/api/useListClockIns";
import usePrivateAxios from "../hooks/usePrivateAxios";

const DashboardPage = () => {
  const { t } = useTranslation();
  const theme = useTheme();
  const { authState } = useContext(authContext);
  const privateAxios = usePrivateAxios();
  const { clockIns, loading } = useListClockIns(privateAxios);
  const { clockIn, clockOut } = useClockInsClockOutsApi(privateAxios);

  return (
    <>
      <Helmet>
        <title>{t("nav.mypage")}</title>
      </Helmet>
      <Container maxWidth="xl">
        <Stack
          direction="row"
          alignItems="center"
          justifyContent="space-between"
          sx={{ mb: 5 }}
        >
          <Stack>
            <Typography variant="h3">{t("dashboard.welcome")}</Typography>
            <Typography variant="subtitle1" sx={{ opacity: 0.8 }}>
              {authState.user?.first_name} {authState.user?.last_name}
            </Typography>
          </Stack>
          <Clock />
        </Stack>
        {loading ? (
          <Loading />
        ) : (
          <Grid justifyContent="space-evenly" container spacing={3}>
            <Grid item xs={12} sm={5}>
              <CustomButton
                title={t("dashboard.clock_in")}
                color={theme.palette.success.contrastText}
                bgcolor={theme.palette.success.light}
                onClick={clockIn}
              />
            </Grid>
            <Grid item xs={12} sm={5}>
              <CustomButton
                title={t("dashboard.clock_out")}
                color={theme.palette.warning.contrastText}
                bgcolor={theme.palette.warning.light}
                onClick={clockOut}
              />
            </Grid>
            <Grid item xs={12}>
              <ClockInsList data={clockIns} />
            </Grid>
          </Grid>
        )}
      </Container>
    </>
  );
};

export default DashboardPage;
