import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import DownloadIcon from '@mui/icons-material/Download';
import {AppBar, Box, IconButton, Toolbar} from "@mui/material";

export default function Index({backUrl, downloadHref}: {
  backUrl: string
  downloadHref: string
}) {
  return (
    <Box>
      <AppBar
        position='absolute'
        color='transparent'
        sx={theme => ({
          color: theme.palette.background.paper,
          boxShadow: 'none',
          '& a': {
            color: 'inherit',
          },
          '& a.Mui-focusVisible, & a:hover': {
            backgroundColor: 'rgba(255, 255, 255, 0.2)',
          },
        })}
      >
        <Toolbar>
          <IconButton href={backUrl}>
            <ArrowBackIcon/>
          </IconButton>

          <Box sx={{flexGrow: 1}}/>

          <IconButton href={downloadHref}
                      title='Download'
                      component='a'
                      download>
            <DownloadIcon/>
          </IconButton>
        </Toolbar>
      </AppBar>
    </Box>
  )
}