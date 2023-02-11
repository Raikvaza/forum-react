import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import './Header.css';

const Header = (props) => {  
const handleMouseMovement = (e) => {
    const x = e.pageX - e.target.offsetLeft
      const y = e.pageY - e.target.offsetTop
    
      e.target.style.setProperty('--x', `${ x }px`)
      e.target.style.setProperty('--y', `${ y }px`)
}
const [auth, setAuth] = useState(false);

useEffect(() => {
if(props.status){
  if(props.status === true){
    setAuth(true)
  } 
}   
})

const handleAuth = () => {
  if (auth){
    return(
      <>
      <Link to="/signin">
          <button className="button" onMouseMove={handleMouseMovement}>Profile</button>
        </Link>
        <Link to="/signup">
          <button className="button" onMouseMove={handleMouseMovement}>Log Out</button>
        </Link>
      </>
    )
  } else {
    return (
    <>
      <Link to="/signin">
          <button className="button" onMouseMove={handleMouseMovement}>Log In</button>
        </Link>
        <Link to="/signup">
          <button className="button" onMouseMove={handleMouseMovement}>Sign Up</button>
        </Link>
      </>
      )
  }
  
}


  return (
    <header className="header-container">
      <div className="nav">
        <Link to="/">
          <button className="button" onMouseMove={handleMouseMovement}>Home</button>
        </Link>
        <Link to="/posts">
          <button className="button" onMouseMove={handleMouseMovement}>Posts</button>
        </Link>
        
        {auth && 
        <>
          <Link to="/likedposts">
            <button className="button" onMouseMove={handleMouseMovement}>Liked Posts</button>
          </Link>
          <Link to="/createpost">
            <button className="button" onMouseMove={handleMouseMovement}>Create Post</button>
          </Link>
        </>
        }
        
        
      </div>
      
      <div className="title" onMouseMove={handleMouseMovement}>
        <h1>FORUM</h1>
      </div>
      <div className="authorization">
      
      {handleAuth()}
      </div>
        
    </header>

  );
};

export default Header;