import { toast } from "react-hot-toast";
import usePrivateAxios from "../usePrivateAxios";

const useEmployees = () => {
  const api = usePrivateAxios();

  const getEmployeeById = async (id: string) => {
    const url = `employees/${id}`;
    try {
      const response = await api.get(url, {});
      return response.data;
    } catch (error) {
      toast.error(`an error occured, ${error}`);
    }
  };

  return { getEmployeeById };
};

export default useEmployees;
