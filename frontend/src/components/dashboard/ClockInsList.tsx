import {
  Card,
  Chip,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@mui/material";
import { ListClockInsResponse } from "../../services/clockins";
import formatTime from "../../utils/formatTime";

type Props = {
  data: ListClockInsResponse[];
};

const ClockInsList = ({ data }: Props) => {
  const calcWorkingHours = (clockIn: ListClockInsResponse) => {
    const clockOutTime = new Date(clockIn.clock_out_time.Time).getTime();
    const clockInTime = new Date(clockIn.clock_in_time).getTime();
    const diffMillis = clockOutTime - clockInTime;
    const diffMinutes = Math.floor(diffMillis / (1000 * 60));
    const hours = Math.floor(diffMinutes / 60);
    const minutes = diffMinutes % 60;
    return `${hours.toString().padStart(2, "0")}:${minutes
      .toString()
      .padStart(2, "0")}`;
  };

  return (
    <>
      <Card>
        <TableContainer>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>ClockIn Time</TableCell>
                <TableCell>ClockOut Time</TableCell>
                <TableCell>Clocked Out</TableCell>
                <TableCell>Working Hours</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {data.map((clockIn) => (
                <TableRow hover key={clockIn.clock_in_id}>
                  <TableCell align="left">
                    {formatTime(clockIn.clock_in_time)}
                  </TableCell>
                  <TableCell align="left">
                    {clockIn.clock_out_time.Valid
                      ? formatTime(clockIn.clock_out_time.Time)
                      : "N/A"}
                  </TableCell>
                  <TableCell align="left">
                    {clockIn.clock_out_time.Valid ? (
                      <Chip color="success" label="Yes" />
                    ) : (
                      <Chip color="error" label="No" />
                    )}
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
