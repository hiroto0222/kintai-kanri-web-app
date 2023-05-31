import { AxiosInstance, AxiosResponse } from "axios";
import { useCallback, useEffect, useState } from "react";
import { toast } from "react-hot-toast";
import { ListClockInsResponse } from "../../services/clockins";

const useListClockIns = (
  privateApi: AxiosInstance,
  employeeId: string | undefined
) => {
  const [loading, setLoading] = useState(true);
  const [clockIns, setClockIns] = useState<ListClockInsResponse[]>([]);
  const [latestClockIn, setLatestClockIn] = useState<ListClockInsResponse>();

  const getClockIns = useCallback(
    async (employeeId: string | undefined) => {
      try {
        const url = `clockins/${employeeId}`;
        const response: AxiosResponse<ListClockInsResponse[]> =
          await privateApi.get(url, { withCredentials: true });
        setLoading(false);
        setClockIns(response.data);
        if (response.data && response.data.length > 0) {
          setLatestClockIn(response.data[0]);
        }
      } catch (error) {
        toast.error(`an error occured, ${error}`);
        setLoading(false);
        setClockIns([]);
      }
    },
    [privateApi]
  );

  useEffect(() => {
    getClockIns(employeeId);
  }, [privateApi, getClockIns, employeeId]);

  return { clockIns, latestClockIn, loading, getClockIns };
};

export default useListClockIns;
