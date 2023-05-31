import { Card, Typography, styled, useTheme } from "@mui/material";

type Props = {
  title: string;
  color: string;
  bgcolor: string;
  onClick: () => Promise<void>;
  disabled: boolean;
};

const StyledCard = styled(Card)(() => ({
  transition: "transform 0.2s ease-out",
  cursor: "pointer",
  "&:hover": {
    transform: "scale(1.05)",
  },
}));

const Button = ({ title, color, bgcolor, onClick, disabled }: Props) => {
  const theme = useTheme();

  return (
    <>
      {disabled ? (
        <Card
          sx={{
            py: 5,
            boxShadow: 0,
            textAlign: "center",
            color,
            bgcolor: theme.palette.grey[400],
          }}
        >
          <Typography variant="h3">{title}</Typography>
        </Card>
      ) : (
        <StyledCard
          sx={{
            py: 5,
            boxShadow: 0,
            textAlign: "center",
            color,
            bgcolor,
          }}
          onClick={onClick}
        >
          <Typography variant="h3">{title}</Typography>
        </StyledCard>
      )}
    </>
  );
};

export default Button;
