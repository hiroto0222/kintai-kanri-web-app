import { Grid, useTheme } from "@mui/material";
import { useContext } from "react";
import { useTranslation } from "react-i18next";
import { authContext } from "../../context/auth";
import { ListClockInsResponse } from "../../services/clockins";
import CustomButton from "./Button";

type Props = {
  clockIn: () => Promise<void>;
  clockOut: () => Promise<void>;
  getClockIns: (employeeId: string | undefined) => Promise<void>;
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
  const { authState } = useContext(authContext);

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
            await getClockIns(authState.user?.id);
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
            await getClockIns(authState.user?.id);
          }}
          disabled={isClockOutDisabled()}
        />
      </Grid>
    </>
  );
};

export default ClockInClockOut;
