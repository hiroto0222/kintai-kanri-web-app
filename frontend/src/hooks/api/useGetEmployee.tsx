import { AxiosInstance, AxiosResponse } from "axios";
import { useEffect, useState } from "react";
import { toast } from "react-hot-toast";
import { Employee } from "../../services/employees";

const useGetEmployee = (
  privateApi: AxiosInstance,
  employeeId: string | undefined
) => {
  const [employee, setEmployee] = useState<Employee>();

  useEffect(() => {
    const getEmployeeById = async () => {
      const url = `employees/${employeeId}`;
      try {
        const response: AxiosResponse<Employee> = await privateApi.get(url, {
          withCredentials: true,
        });
        setEmployee(response.data);
      } catch (error) {
        toast.error(`an error occured, ${error}`);
      }
    };

    getEmployeeById();
  }, [privateApi, employeeId]);

  return { employee };
};

export default useGetEmployee;
