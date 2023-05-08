import usePrivateAxios from "../usePrivateAxios";

const useEmployees = () => {
  const api = usePrivateAxios();

  const getEmployeeById = async (id: string) => {
    const url = `employees/${id}`;
    try {
      const response = await api.get(url, {});
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log(error);
    }
  };

  return { getEmployeeById };
};

export default useEmployees;
