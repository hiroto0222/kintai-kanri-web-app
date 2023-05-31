import {
  Card,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@mui/material";
import { useTranslation } from "react-i18next";
import { ListClockInsResponse } from "../../services/clockins";
import { calcWorkingHours, formatDate, formatTime } from "../../utils";

type Props = {
  data: ListClockInsResponse[];
};

const ClockInsList = ({ data }: Props) => {
  const { t } = useTranslation();

  return (
    <>
      <Card>
        <TableContainer>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>{t("clockinsTable.date")}</TableCell>
                <TableCell>{t("clockinsTable.clock_in_time")}</TableCell>
                <TableCell>{t("clockinsTable.clock_out_time")}</TableCell>
                <TableCell>{t("clockinsTable.working_hours")}</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {data &&
                data.map((clockIn) => (
                  <TableRow hover key={clockIn.clock_in_id}>
                    <TableCell align="left">
                      {formatDate(clockIn.clock_in_time)}
                    </TableCell>
                    <TableCell align="left">
                      {formatTime(clockIn.clock_in_time)}
                    </TableCell>
                    <TableCell align="left">
                      {clockIn.clock_out_time.Valid
                        ? formatTime(clockIn.clock_out_time.Time)
                        : "N/A"}
                    </TableCell>
                    <TableCell align="left">
                      {clockIn.clock_out_time.Valid
                        ? calcWorkingHours(clockIn)
                        : "N/A"}
                    </TableCell>
                  </TableRow>
                ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Card>
    </>
  );
};

export default ClockInsList;
