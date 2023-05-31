import { Card, Typography, styled } from "@mui/material";

type Props = {
  title: string;
  color: string;
  bgcolor: string;
  onClick: () => Promise<void>;
};

const StyledCard = styled(Card)(() => ({
  transition: "transform 0.2s ease-out",
  cursor: "pointer",
  "&:hover": {
    transform: "scale(1.05)",
  },
}));

const Button = ({ title, color, bgcolor, onClick }: Props) => {
  return (
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
      {/* <Typography variant="subtitle2" sx={{ opacity: 0.8 }}>
        {title}
      </Typography> */}
    </StyledCard>
  );
};

export default Button;
