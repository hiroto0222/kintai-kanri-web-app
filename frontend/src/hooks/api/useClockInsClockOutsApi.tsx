import { AxiosInstance } from "axios";
import { toast } from "react-hot-toast";

const useClockInsClockOutsApi = (privateApi: AxiosInstance) => {
  const clockIn = async (employeeID: string) => {
    const url = "clockins";
    try {
      await privateApi.post(
        url,
        {
          employee_id: employeeID,
        },
        {
          withCredentials: true,
        }
      );
      toast.success("successfully clocked in!");
    } catch (error) {
      toast.error(`an error occured, ${error}`);
    }
  };

  const clockOut = async (employeeID: string) => {
    const url = "clockouts";
    try {
      await privateApi.post(
        url,
        { employee_id: employeeID },
        { withCredentials: true }
      );
      toast.success("successfully clocked out!");
    } catch (error) {
      toast.error(`an error occured, ${error}`);
    }
  };

  return { clockIn, clockOut };
};

export default useClockInsClockOutsApi;
