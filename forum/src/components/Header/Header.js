import React from "react";
import { Link } from "react-router-dom";
import './Header.css';
const Header = () => {  
const handleMouseMovement = (e) => {
    const x = e.pageX - e.target.offsetLeft
      const y = e.pageY - e.target.offsetTop
    
      e.target.style.setProperty('--x', `${ x }px`)
      e.target.style.setProperty('--y', `${ y }px`)
}

  return (
    <header className="header-container">
      <div className="nav">
        <Link to="/home">
          <button className="button" onMouseMove={handleMouseMovement}>Home</button>
        </Link>
        <Link to="/posts">
          <button className="button" onMouseMove={handleMouseMovement}>Posts</button>
        </Link>
        <Link to="/likedposts">
          <button className="button" onMouseMove={handleMouseMovement}>Liked Posts</button>
        </Link>
        <Link to="/createpost">
          <button className="button" onMouseMove={handleMouseMovement}>Create Post</button>
        </Link>
      </div>
      
      <div className="title" onMouseMove={handleMouseMovement}>
        <h1>FORUM</h1>
      </div>
      <div className="authorization">
      <Link to="/signin">
          <button className="button" onMouseMove={handleMouseMovement}>Log In</button>
        </Link>
        <Link to="/signup">
          <button className="button" onMouseMove={handleMouseMovement}>Sign Up</button>
        </Link>
      </div>
        
    </header>

  );
};

export default Header;