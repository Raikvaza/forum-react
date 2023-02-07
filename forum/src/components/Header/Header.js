import React from "react";
import { Link } from "react-router-dom";
import { withStyles, makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";

const useStyles = makeStyles(theme => ({
  root: {
    background: "black",
    display: "flex",
    justifyContent: "space-between",
    alignItems: "center",
    padding: "1rem",
    [theme.breakpoints.down("sm")]: {
      flexDirection: "column",
      alignItems: "center"
    }
  },
  button: {
    color: "white",
    border: "1px solid white",
    borderRadius: "20px",
    margin: "1rem",
    [theme.breakpoints.down("sm")]: {
      margin: "0.5rem 0"
    }
  }
}));

const StyledButton = withStyles({
  root: {
    background: "transparent",
    border: "1px solid white",
    borderRadius: "20px",
    color: "white",
    "&:hover": {
      background: "white",
      color: "black"
    }
  }
})(Button);

const Header = () => {
  const classes = useStyles();

  return (
    <header className={classes.root}>
      <Link to="/">
        <StyledButton className={classes.button}>Home</StyledButton>
      </Link>
      <Link to="/posts">
        <StyledButton className={classes.button}>Posts</StyledButton>
      </Link>
      <Link to="/likedposts">
        <StyledButton className={classes.button}>Liked Posts</StyledButton>
      </Link>
      <Link to="/createpost">
        <StyledButton className={classes.button}>Create Post</StyledButton>
      </Link>
    </header>
  );
};

export default Header;