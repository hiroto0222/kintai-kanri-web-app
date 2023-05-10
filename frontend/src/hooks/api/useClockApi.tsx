import { toast } from "react-hot-toast";
import usePrivateAxios from "../usePrivateAxios";

const useClockApi = () => {
  const api = usePrivateAxios();

  const clockIn = async (employeeID: string) => {
    const url = "clockins";
    try {
      await api.post(
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

  return { clockIn };
};

export default useClockApi;
