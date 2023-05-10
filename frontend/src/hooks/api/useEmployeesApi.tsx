import { AxiosInstance, AxiosResponse } from "axios";
import { toast } from "react-hot-toast";
import { Employee } from "../../services/employees";

const useEmployees = (privateApi: AxiosInstance) => {
  const getEmployeeById = async (id: string) => {
    const url = `employees/${id}`;
    try {
      const response: AxiosResponse<Employee> = await privateApi.get(url, {
        withCredentials: true,
      });
      return response.data;
    } catch (error) {
      toast.error(`an error occured, ${error}`);
    }
  };

  return { getEmployeeById };
};

export default useEmployees;
