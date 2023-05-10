import { AxiosInstance, AxiosResponse } from "axios";
import { useContext, useEffect, useState } from "react";
import { toast } from "react-hot-toast";
import { authContext } from "../../context/auth";
import { ListClockInsResponse } from "../../services/clockins";

const useListClockIns = (privateApi: AxiosInstance) => {
  const [loading, setLoading] = useState(true);
  const [clockIns, setClockIns] = useState<ListClockInsResponse[]>([]);
  const { authState } = useContext(authContext);

  useEffect(() => {
    const listClockIns = async (employeeID: string) => {
      const url = `clockins/${employeeID}`;
      try {
        const response: AxiosResponse<ListClockInsResponse[]> =
          await privateApi.get(url, { withCredentials: true });
        setLoading(false);
        setClockIns(response.data);
      } catch (error) {
        toast.error(`an error occured, ${error}`);
        setLoading(false);
        setClockIns([]);
      }
    };
    if (authState.user) {
      listClockIns(authState.user.id);
    }
  }, [authState, privateApi]);

  return { clockIns, loading };
};

export default useListClockIns;
