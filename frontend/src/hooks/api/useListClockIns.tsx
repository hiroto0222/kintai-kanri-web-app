import { AxiosInstance, AxiosResponse } from "axios";
import { useCallback, useContext, useEffect, useState } from "react";
import { toast } from "react-hot-toast";
import { authContext } from "../../context/auth";
import { ListClockInsResponse } from "../../services/clockins";

const useListClockIns = (privateApi: AxiosInstance) => {
  const [loading, setLoading] = useState(true);
  const [clockIns, setClockIns] = useState<ListClockInsResponse[]>([]);
  const [latestClockIn, setLatestClockIn] = useState<ListClockInsResponse>();
  const { authState } = useContext(authContext);

  const getClockIns = useCallback(async () => {
    try {
      const url = `clockins/${authState.user?.id}`;
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
  }, [privateApi, authState]);

  useEffect(() => {
    getClockIns();
  }, [authState, privateApi, getClockIns]);

  return { clockIns, latestClockIn, loading, getClockIns };
};

export default useListClockIns;
