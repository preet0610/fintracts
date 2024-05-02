import { Routes, Route } from "react-router-dom";
import SignUp from "./pages/signup";
import Login from "./pages/login";
import EmployerSignUp from "./pages/signup-employer";
import BankSignUp from "./pages/signup-bank";
import CentralBankSignUp from "./pages/signup-centralbank";
import ForexSignUp from "./pages/signup-forexbank";
import EmployeeSignUp from "./pages/signup-employee";
import Home from "./pages/home";
import CreateContract from "./pages/createContract";
import AgreeContract from "./pages/agreeContract";
import RevokedContracts from "./pages/revokedContracts";
import AddExchange from "./pages/addExchange";
import BankPending from "./pages/bankPending";
import CreateBankAcc from "./pages/createBankAcc";
import ViewBalance from "./pages/veiwBalance";
import CentralBankPending from "./pages/centralBankPending";
import ForexPending from "./pages/forexPending";
import "./App.css";

const App = () => {
  const login = window.localStorage.getItem("isLoggedIn");
  return (
    <Routes>
      <Route path="/" element={login ? <Home /> : <SignUp />} />
      <Route path="/signup" element={<SignUp />} />
      <Route path="/login" element={<Login />} />
      <Route path="/signup/employer" element={<EmployerSignUp />} />
      <Route path="/signup/employee" element={<EmployeeSignUp />} />
      <Route path="/signup/bank" element={<BankSignUp />} />
      <Route path="/signup/centralbank" element={<CentralBankSignUp />} />
      <Route path="/signup/forex" element={<ForexSignUp />} />
      <Route path="/home" element={<Home />} />
      <Route path="/create-contract" element={<CreateContract />} />
      <Route path="/agree-to-contract" element={<AgreeContract />} />
      <Route path="/revoked-contracts" element={<RevokedContracts />} />
      <Route path="/add-exchange-rate" element={<AddExchange />} />
      <Route path="/bank-pending-txns" element={<BankPending />} />
      <Route path="/create-bank-account" element={<CreateBankAcc />} />
      <Route path="/veiw-balance" element={<ViewBalance />} />
      <Route path="/forex-pending-txns" element={<ForexPending />} />
      <Route
        path="/central-bank-pending-txns"
        element={<CentralBankPending />}
      />
    </Routes>
  );
};

export default App;
