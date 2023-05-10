import { AxiosInstance } from "axios";
import { useContext } from "react";
import { toast } from "react-hot-toast";
import { authContext } from "../../context/auth";

const useClockInsClockOutsApi = (privateApi: AxiosInstance) => {
  const { authState } = useContext(authContext);

  const clockIn = async () => {
    const url = "clockins";
    try {
      await privateApi.post(
        url,
        {
          employee_id: authState.user?.id,
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

  const clockOut = async () => {
    const url = "clockouts";
    try {
      await privateApi.post(
        url,
        { employee_id: authState.user?.id },
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
