import { Grid, useTheme } from "@mui/material";
import { useTranslation } from "react-i18next";
import { ListClockInsResponse } from "../../services/clockins";
import CustomButton from "./Button";

type Props = {
  clockIn: () => Promise<void>;
  clockOut: () => Promise<void>;
  getClockIns: () => Promise<void>;
  latestClockIn: ListClockInsResponse | undefined;
};

const ClockInClockOut = ({
  clockIn,
  clockOut,
  getClockIns,
  latestClockIn,
}: Props) => {
  const { t } = useTranslation();
  const theme = useTheme();

  // まだ退勤打刻していない場合はclockInをdisabled
  const isClockInDisabled = () => {
    return latestClockIn === undefined
      ? false
      : !latestClockIn.clock_out_id.Valid;
  };

  // まだ出勤打刻していない場合はclockOutをdisabled
  const isClockOutDisabled = () => {
    return latestClockIn === undefined
      ? true
      : latestClockIn.clock_out_id.Valid;
  };

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
          disabled={isClockInDisabled()}
        />
      </Grid>
      <Grid item xs={12} sm={5}>
        <CustomButton
          title={t("dashboard.clock_out")}
          color={theme.palette.warning.contrastText}
          bgcolor={theme.palette.error.light}
          onClick={async () => {
            await clockOut();
            await getClockIns();
          }}
          disabled={isClockOutDisabled()}
        />
      </Grid>
    </>
  );
};

export default ClockInClockOut;
