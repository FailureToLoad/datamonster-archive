import { UserButton } from "@clerk/clerk-react";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import {
  Archive,
  Person,
  IconProps,
  HourglassMedium,
} from "@phosphor-icons/react";
import { Link, Outlet, useLocation } from "react-router-dom";

function LeftNav() {
  const { pathname } = useLocation();
  const timelineKey = "timeline";
  const populationKey = "population";
  const storageKey = "storage";
  const getProps = (active: boolean): IconProps => {
    const props: IconProps = {
      size: 32,
    };
    if (active) {
      props.weight = "fill";
      props.className = "text-primary";
    }
    return props;
  };
  return (
    <div className="h-screen absolute grid top-0 left-0 p-4">
      <div className="flex flex-col gap-4">
        <TooltipProvider delayDuration={400}>
          <Tooltip>
            <TooltipTrigger asChild>
              <Link to={timelineKey} color="foreground">
                <HourglassMedium
                  {...getProps(pathname.includes(timelineKey))}
                />
              </Link>
            </TooltipTrigger>
            <TooltipContent side="right">
              <p>Timeline</p>
            </TooltipContent>
          </Tooltip>
        </TooltipProvider>
        <TooltipProvider delayDuration={400}>
          <Tooltip>
            <TooltipTrigger asChild>
              <Link to={populationKey} color="foreground">
                <Person {...getProps(pathname.includes(populationKey))} />
              </Link>
            </TooltipTrigger>
            <TooltipContent side="right">
              <p>Population</p>
            </TooltipContent>
          </Tooltip>
        </TooltipProvider>
        <TooltipProvider delayDuration={400}>
          <Tooltip>
            <TooltipTrigger asChild>
              <Link to={storageKey} color="foreground">
                <Archive {...getProps(pathname.includes(storageKey))} />
              </Link>
            </TooltipTrigger>
            <TooltipContent side="right">
              <p>Storage</p>
            </TooltipContent>
          </Tooltip>
        </TooltipProvider>
      </div>

      <div className="flex self-end justify-self-end place-self-end">
        <UserButton />
      </div>
    </div>
  );
}

export default function SettlementPage() {
  return (
    <>
      <LeftNav />
      <div className="flex h-screen w-full flex-col justify-center overflow-auto">
        <div className="p-16 flex flex-1 justify-center">
          <Outlet />
        </div>
      </div>
    </>
  );
}
