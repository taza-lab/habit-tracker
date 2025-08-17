import { Typography, Box } from "@mui/material";

type PageTitleProps = {
    title: string;
};

const PageTitle = ({ title }: PageTitleProps) => (
    <Box sx={{ mb: 3 }}>
        <Typography variant="h4" component="h2" fontWeight="bold">
            {title}
        </Typography>
    </Box>
);

export default PageTitle;