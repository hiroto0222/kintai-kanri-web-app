import { AxiosInstance, AxiosResponse } from "axios";
import { useContext, useEffect, useState } from "react";
import { toast } from "react-hot-toast";
import { authContext } from "../../context/auth";
import { Employee } from "../../services/employees";

const useListEmployees = (privateApi: AxiosInstance) => {
  const [loading, setLoading] = useState(true);
  const [employees, setEmployees] = useState<Employee[]>([]);
  const { authState } = useContext(authContext);

  useEffect(() => {
    const listEmployees = async () => {
      const url = `employees`;
      try {
        const response: AxiosResponse<Employee[]> = await privateApi.get(url, {
          withCredentials: true,
          params: {
            page_id: 1,
            page_size: 20, // TODO: add pagination
          },
        });
        setLoading(false);
        setEmployees(response.data);
      } catch (error) {
        toast.error(`an error occured, ${error}`);
        setLoading(false);
        setEmployees([]);
      }
    };
    listEmployees();
  }, [authState, privateApi]);

  return { employees, loading };
};

export default useListEmployees;
