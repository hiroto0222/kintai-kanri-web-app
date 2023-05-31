import { ListClockInsResponse } from "../services/clockins";

export const formatTime = (s: string) => {
  return new Date(s).toLocaleTimeString();
};

export const formatDate = (s: string) => {
  return new Date(s).toLocaleDateString();
};

export const calcWorkingHours = (clockIn: ListClockInsResponse) => {
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
