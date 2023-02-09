import React from 'react';
import { BrowserRouter, Route, Routes} from "react-router-dom";
import ReactDOM from 'react-dom';

import LoginPage from './routes/Login-Page/LoginPage';
import SignUpPage from './routes/Sign-Up-Page/SignUpPage';
import HomePage from './routes/Home-Page/HomePage';
import CreatePost from './routes/Create-Post/CreatePost';

// const AuthContext = createContext();

// function ProvideAuth({ children }) {
//   const [isAuth, setIsAuth] = useState(false);

//   return (
//     <AuthContext.Provider value={{ isAuth, setIsAuth }}>
//       {children}
//     </AuthContext.Provider>
//   );
// }

function App () {

  // const [isAuth, setAuth] = useState(false);
  // const [User, setUser] = useState('a')
  return (
    //<AuthContext.Provider value={{isAuth, setAuth, User, setUser}}>
    <BrowserRouter>
      <Routes>
        <Route exact path="/" element={<HomePage/>} />
        <Route path="/signup" element={<SignUpPage/>} />
        <Route path="/signin" element={<LoginPage/>} />
        <Route path="/createpost" element={<CreatePost/>} />
      </Routes>
    </BrowserRouter>
  //</AuthContext.Provider>
  )
};
export default App;
const container = document.getElementById('root');
ReactDOM.render(<App />, container);

//const root = ReactDOM.createRoot(document.getElementById('root'));
// root.render(
//   <React.StrictMode>
//     <App />
//     {/* <RouterProvider router={router} /> */}
//   </React.StrictMode>
// );