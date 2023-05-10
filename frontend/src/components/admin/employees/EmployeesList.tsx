import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@mui/material";
import { useNavigate } from "react-router-dom";
import { Employee } from "../../../services/employees";

type Props = {
  data: Employee[];
};

const EmployeesList = ({ data }: Props) => {
  const navigate = useNavigate();

  const handleClick = (employeeID: string) => {
    console.log(employeeID);
    // navigate(`/dashboard/employees/${employeeID}`);
  };

  return (
    <>
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>First Name</TableCell>
              <TableCell>Last Name</TableCell>
              <TableCell>Email</TableCell>
              <TableCell>Phone</TableCell>
              <TableCell>Address</TableCell>
              <TableCell>Role</TableCell>
              <TableCell>Joined</TableCell>
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
                <TableCell>{employee.first_name}</TableCell>
                <TableCell>{employee.last_name}</TableCell>
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
