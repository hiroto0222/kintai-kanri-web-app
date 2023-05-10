import { Grid, TextField } from "@mui/material";
import { Control, Controller } from "react-hook-form";

export type LoginFormProps = {
  Email: string;
  Password: string;
};

type Props = {
  control: Control<LoginFormProps>;
};

type formData = {
  itemProps: { xs: number; md?: number };
  name: keyof LoginFormProps;
  label: string;
  required: boolean;
  type: React.HTMLInputTypeAttribute;
  defaultValue: string;
  rules?: any;
};

const loginFormData: formData[] = [
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
    name: "Password",
    label: "Password",
    required: true,
    type: "password",
    defaultValue: "",
  },
];

const LoginForm: React.FC<Props> = ({ control }) => {
  return (
    <Grid container spacing={2} marginY={2}>
      {loginFormData.map((item) => (
        <Grid item {...item.itemProps} key={item.label}>
          <Controller
            name={item.name}
            defaultValue={item.defaultValue}
            control={control}
            rules={{
              ...item.rules,
            }}
            render={({
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
            )}
          />
        </Grid>
      ))}
    </Grid>
  );
};

export default LoginForm;
