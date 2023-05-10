export type CreateClockInResponse = {
  id: number;
  clocked_out: boolean;
  clock_in_time: string;
  employee_id: string;
};

export type ListClockInsResponse = {
  employee_id: string;
  clock_in_id: number;
  clock_in_time: string;
  clock_out_id: {
    Int32: number; // sql.NullInt32
    Valid: boolean;
  };
  clock_out_time: {
    Time: string; // sql.NullTime
    Valid: boolean;
  };
};
