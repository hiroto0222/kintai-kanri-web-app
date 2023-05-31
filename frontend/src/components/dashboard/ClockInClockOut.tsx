import { Grid, useTheme } from "@mui/material";
import { useTranslation } from "react-i18next";
import { ListClockInsResponse } from "../../services/clockins";
import CustomButton from "./Button";

type Props = {
  clockIn: () => Promise<void>;
  clockOut: () => Promise<void>;
  getClockIns: () => Promise<void>;
  clockIns: ListClockInsResponse[];
};

const ClockInClockOut = ({
  clockIn,
  clockOut,
  getClockIns,
  clockIns,
}: Props) => {
  const { t } = useTranslation();
  const theme = useTheme();

  console.log(clockIns);

  return (
    <>
      <Grid item xs={12} sm={5}>
        <CustomButton
          title={t("dashboard.clock_in")}
          color={theme.palette.success.contrastText}
          bgcolor={theme.palette.success.light}
          onClick={async () => {
            await clockIn();
            await getClockIns();
          }}
        />
      </Grid>
      <Grid item xs={12} sm={5}>
        <CustomButton
          title={t("dashboard.clock_out")}
          color={theme.palette.warning.contrastText}
          bgcolor={theme.palette.warning.light}
          onClick={async () => {
            await clockOut();
            await getClockIns();
          }}
        />
      </Grid>
    </>
  );
};

export default ClockInClockOut;
