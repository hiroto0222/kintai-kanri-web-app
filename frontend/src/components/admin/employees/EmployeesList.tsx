import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@mui/material";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";
import { Employee } from "../../../services/employees";

type Props = {
  data: Employee[];
};

const EmployeesList = ({ data }: Props) => {
  const { t } = useTranslation();
  const navigate = useNavigate();

  const handleClick = (employeeID: string) => {
    navigate(`/dashboard/employees/${employeeID}`);
  };

  return (
    <>
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>{t("employeesTable.last_name")}</TableCell>
              <TableCell>{t("employeesTable.first_name")}</TableCell>
              <TableCell>{t("employeesTable.email")}</TableCell>
              <TableCell>{t("employeesTable.phone")}</TableCell>
              <TableCell>{t("employeesTable.address")}</TableCell>
              <TableCell>{t("employeesTable.role")}</TableCell>
              <TableCell>{t("employeesTable.joined")}</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {data.map((employee) => (
              <TableRow
                hover
                key={employee.id}
                onClick={() => handleClick(employee.id)}
                sx={{
                  cursor: "pointer",
                }}
              >
                <TableCell>{employee.last_name}</TableCell>
                <TableCell>{employee.first_name}</TableCell>
                <TableCell>{employee.email}</TableCell>
                <TableCell>{employee.phone}</TableCell>
                <TableCell>{employee.address}</TableCell>
                <TableCell>N/A</TableCell>
                <TableCell>
                  {new Date(employee.created_at).toLocaleDateString()}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
};

export default EmployeesList;
