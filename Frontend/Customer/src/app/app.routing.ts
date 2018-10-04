import { CustomerComponent } from "./customer/customer.component";
import { CreateComponent } from "./create/create.component";

export const AvailableRoutes: any = [
    { path: "", component: CustomerComponent },
    { path: "create", component: CreateComponent }
];