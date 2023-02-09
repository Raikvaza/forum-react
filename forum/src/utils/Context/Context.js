import React, { createContext, useContext, useEffect, useState } from 'react';

// YourContextProvider = (props)=>{

//     yourContext = React.createContext(null);
  
//     value = { shared, state, and, functions }
//       return (
//          <YourContext.Provider value={value}>
//            {props.children}
//          </YourContext.Provider>
//       )
//   }
const LoginContext = createContext();

function ProvideLogin({ children }) {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [GlobalUsername, setGlobalUsername] = useState('')
    return (
      <LoginContext.Provider value={{ isLoggedIn, setIsLoggedIn, GlobalUsername, setGlobalUsername }}>
        {children}
      </LoginContext.Provider>
    );
  }
  export {ProvideLogin, LoginContext}