import { Checkbox, FormControlLabel, Grid, TextField } from "@mui/material";
import { Control, Controller } from "react-hook-form";

export type RegisterFormProps = {
  FirstName: string;
  LastName: string;
  Email: string;
  Phone: string;
  Address: string;
  RoleID?: number;
  IsAdmin: boolean;
  Password: string;
};

type Props = {
  control: Control<RegisterFormProps>;
};

type formData = {
  itemProps: { xs: number; md?: number };
  name: keyof RegisterFormProps;
  label: string;
  required: boolean;
  type: React.HTMLInputTypeAttribute;
  defaultValue: string | number | boolean | undefined;
  rules?: any;
};

const registerFormData: formData[] = [
  {
    itemProps: { xs: 12, md: 6 },
    name: "FirstName",
    label: "First Name",
    required: true,
    type: "text",
    defaultValue: "",
    rules: {
      maxLength: {
        value: 50,
      },
    },
  },
  {
    itemProps: { xs: 12, md: 6 },
    name: "LastName",
    label: "Last Name",
    required: true,
    type: "text",
    defaultValue: "",
    rules: {
      maxLength: {
        value: 50,
      },
    },
  },
  {
    itemProps: { xs: 12 },
    name: "Email",
    label: "Email",
    required: true,
    type: "email",
    defaultValue: "",
    rules: {
      maxLength: {
        value: 255,
      },
    },
  },
  {
    itemProps: { xs: 12 },
    name: "Phone",
    label: "Phone Number",
    required: true,
    type: "text",
    defaultValue: "",
    rules: {
      maxLength: {
        value: 20,
      },
    },
  },
  {
    itemProps: { xs: 12 },
    name: "Address",
    label: "Address",
    required: true,
    type: "text",
    defaultValue: "",
    rules: {
      maxLength: {
        value: 255,
      },
    },
  },
  {
    itemProps: { xs: 12 },
    name: "Password",
    label: "Password",
    required: true,
    type: "password",
    defaultValue: "",
    rules: {
      minLength: {
        value: 6,
      },
    },
  },
  {
    itemProps: { xs: 6 },
    name: "RoleID",
    label: "Role ID",
    required: false,
    type: "number",
    defaultValue: undefined,
  },
  {
    itemProps: { xs: 6 },
    name: "IsAdmin",
    label: "Admin",
    required: false,
    type: "checkbox",
    defaultValue: false,
  },
];

const RegisterForm: React.FC<Props> = ({ control }) => {
  return (
    <Grid container spacing={2} marginY={2}>
      {registerFormData.map((item) => (
        <Grid item {...item.itemProps} key={item.label}>
          <Controller
            name={item.name}
            defaultValue={item.defaultValue}
            control={control}
            rules={{
              ...item.rules,
            }}
            render={
              item.type === "checkbox"
                ? ({ field: { onChange, onBlur, value } }) => (
                    <FormControlLabel
                      control={
                        <Checkbox
                          value={value}
                          onChange={onChange}
                          onBlur={onBlur}
                          required={item.required}
                        />
                      }
                      label={item.label}
                    />
                  )
                : ({
                    field: { onChange, onBlur, value },
                    fieldState: { error },
                  }) => (
                    <TextField
                      label={item.label}
                      type={item.type}
                      required={item.required}
                      fullWidth
                      value={value}
                      onChange={onChange}
                      onBlur={onBlur}
                      error={Boolean(error)}
                      helperText={error?.message}
                    />
                  )
            }
          />
        </Grid>
      ))}
    </Grid>
  );
};

export default RegisterForm;
