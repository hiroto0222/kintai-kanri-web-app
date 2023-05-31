import { Container, Grid, Stack, Typography } from "@mui/material";
import { useContext } from "react";
import { Helmet } from "react-helmet-async";
import { useTranslation } from "react-i18next";
import Clock from "../components/dashboard/Clock";
import ClockInClockOut from "../components/dashboard/ClockInClockOut";
import ClockInsList from "../components/dashboard/ClockInsList";
import Loading from "../components/dashboard/Loading";
import { authContext } from "../context/auth";
import useClockInsClockOutsApi from "../hooks/api/useClockInsClockOutsApi";
import useListClockIns from "../hooks/api/useListClockIns";
import usePrivateAxios from "../hooks/usePrivateAxios";

const DashboardPage = () => {
  const { t } = useTranslation();
  const { authState } = useContext(authContext);
  const privateAxios = usePrivateAxios();
  const { clockIns, latestClockIn, getClockIns, loading } =
    useListClockIns(privateAxios);
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
              {authState.user?.last_name} {authState.user?.first_name}
            </Typography>
          </Stack>
          <Clock />
        </Stack>
        {loading ? (
          <Loading />
        ) : (
          <Grid justifyContent="space-evenly" container spacing={3}>
            <ClockInClockOut
              clockIn={clockIn}
              clockOut={clockOut}
              latestClockIn={latestClockIn}
              getClockIns={getClockIns}
            />
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
