import * as React from "react";

import type { SVGProps } from "../global";
import { StyledSVG } from "../global";

const ExperimentIcon = ({ size, ...props }: SVGProps) => (
  <StyledSVG
    size={size}
    viewBox="0 0 24 24"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
    {...props}
  >
    <path
      fillRule="evenodd"
      clipRule="evenodd"
      d="M12.63 6.76594V9.31543L17.4042 14.9034C18.031 15.6374 18.1768 16.6729 17.777 17.5539C17.3775 18.4352 16.5059 19 15.5461 19H2.4539C1.49414 19 0.622504 18.4352 0.222972 17.5539C-0.176752 16.6729 -0.0310314 15.6374 0.595817 14.9034L5.37004 9.31543V3.23679C5.438 3.24512 5.50674 3.24939 5.57605 3.24939H12.424C12.4933 3.24939 12.562 3.24512 12.63 3.23679V5.15947V6.76594ZM13.453 1.03918C13.453 0.763684 13.3449 0.499429 13.1518 0.304388C12.959 0.109735 12.6973 0 12.424 0H5.57605C5.30265 0 5.04097 0.109735 4.84821 0.304388C4.65507 0.499429 4.54698 0.763684 4.54698 1.03918V1.52892C4.54698 1.80481 4.65545 2.06925 4.84898 2.2643C5.04058 2.45817 5.30227 2.56791 5.57605 2.56791H12.424C12.6977 2.56791 12.9594 2.45817 13.1524 2.26274C13.3445 2.06925 13.453 1.80481 13.453 1.52892V1.03918Z"
      fill="#4FBFB5"
    />
  </StyledSVG>
);

export default ExperimentIcon;
