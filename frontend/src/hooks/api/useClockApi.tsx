import usePrivateAxios from "../usePrivateAxios";

const useClockApi = () => {
  const api = usePrivateAxios();

  const clockIn = async (employeeID: string) => {
    const url = "clockins";
    try {
      const response = await api.post(
        url,
        {
          employee_id: employeeID,
        },
        {
          withCredentials: true,
        }
      );
      console.log(response.data);
    } catch (error) {
      console.log(error);
    }
  };

  return { clockIn };
};

export default useClockApi;
