// SPDX-License-Identifier: MIT
// OpenZeppelin Contracts v4.3.2 (token/ERC20/ERC20.sol)
pragma solidity >=0.4.22 <0.9.0;

/**
 * @dev Interface of the ERC20 standard as defined in the EIP.
 */
interface IERC20v2 {
    /**
     * @dev Returns the amount of tokens in existence.
     */
    function totalSupply() external view returns (uint256);

    /**
     * @dev Returns the amount of tokens owned by `account`.
     */
    function balanceOf(address account) external view returns (uint256);

    /**
     * @dev Moves `amount` tokens from the caller's account to `recipient`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transfer(
        address recipient,
        uint256 amount
    ) external returns (bool);

    /**
     * @dev Returns the remaining number of tokens that `spender` will be
     * allowed to spend on behalf of `owner` through {transferFrom}. This is
     * zero by default.
     *
     * This value changes when {approve} or {transferFrom} are called.
     */
    function allowance(
        address owner,
        address spender
    ) external view returns (uint256);

    /**
     * @dev Sets `amount` as the allowance of `spender` over the caller's tokens.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * IMPORTANT: Beware that changing an allowance with this method brings the risk
     * that someone may use both the old and the new allowance by unfortunate
     * transaction ordering. One possible solution to mitigate this race
     * condition is to first reduce the spender's allowance to 0 and set the
     * desired value afterwards:
     * https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
     *
     * Emits an {Approval} event.
     */
    function approve(address spender, uint256 amount) external returns (bool);

    /**
     * @dev Moves `amount` tokens from `sender` to `recipient` using the
     * allowance mechanism. `amount` is then deducted from the caller's
     * allowance.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) external returns (bool);

    /**
     * @dev Emitted when `value` tokens are moved from one account (`from`) to
     * another (`to`).
     *
     * Note that `value` may be zero.
     */
    event Transfer(address indexed from, address indexed to, uint256 value);

    /**
     * @dev Emitted when the allowance of a `spender` for an `owner` is set by
     * a call to {approve}. `value` is the new allowance.
     */
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );
}

library SafeMath {
    /**
     * @dev Returns the addition of two unsigned integers, with an overflow flag.
     *
     * _Available since v3.4._
     */
    function tryAdd(
        uint256 a,
        uint256 b
    ) internal pure returns (bool, uint256) {
        // unchecked {
        uint256 c = a + b;
        if (c < a) return (false, 0);
        return (true, c);
        // }
    }

    /**
     * @dev Returns the substraction of two unsigned integers, with an overflow flag.
     *
     * _Available since v3.4._
     */
    function trySub(
        uint256 a,
        uint256 b
    ) internal pure returns (bool, uint256) {
        // unchecked {
        if (b > a) return (false, 0);
        return (true, a - b);
        //  }
    }

    /**
     * @dev Returns the multiplication of two unsigned integers, with an overflow flag.
     *
     * _Available since v3.4._
     */
    function tryMul(
        uint256 a,
        uint256 b
    ) internal pure returns (bool, uint256) {
        // unchecked {
        // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
        // benefit is lost if 'b' is also tested.
        // See: https://github.com/OpenZeppelin/openzeppelin-contracts/pull/522
        if (a == 0) return (true, 0);
        uint256 c = a * b;
        if (c / a != b) return (false, 0);
        return (true, c);
        // }
    }

    /**
     * @dev Returns the division of two unsigned integers, with a division by zero flag.
     *
     * _Available since v3.4._
     */
    function tryDiv(
        uint256 a,
        uint256 b
    ) internal pure returns (bool, uint256) {
        // unchecked {
        if (b == 0) return (false, 0);
        return (true, a / b);
        // }
    }

    /**
     * @dev Returns the remainder of dividing two unsigned integers, with a division by zero flag.
     *
     * _Available since v3.4._
     */
    function tryMod(
        uint256 a,
        uint256 b
    ) internal pure returns (bool, uint256) {
        // unchecked {
        if (b == 0) return (false, 0);
        return (true, a % b);
        //  }
    }

    /**
     * @dev Returns the addition of two unsigned integers, reverting on
     * overflow.
     *
     * Counterpart to Solidity's `+` operator.
     *
     * Requirements:
     *
     * - Addition cannot overflow.
     */
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        return a + b;
    }

    /**
     * @dev Returns the subtraction of two unsigned integers, reverting on
     * overflow (when the result is negative).
     *
     * Counterpart to Solidity's `-` operator.
     *
     * Requirements:
     *
     * - Subtraction cannot overflow.
     */
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        return a - b;
    }

    /**
     * @dev Returns the multiplication of two unsigned integers, reverting on
     * overflow.
     *
     * Counterpart to Solidity's `*` operator.
     *
     * Requirements:
     *
     * - Multiplication cannot overflow.
     */
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        return a * b;
    }

    /**
     * @dev Returns the integer division of two unsigned integers, reverting on
     * division by zero. The result is rounded towards zero.
     *
     * Counterpart to Solidity's `/` operator.
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        return a / b;
    }

    /**
     * @dev Returns the remainder of dividing two unsigned integers. (unsigned integer modulo),
     * reverting when dividing by zero.
     *
     * Counterpart to Solidity's `%` operator. This function uses a `revert`
     * opcode (which leaves remaining gas untouched) while Solidity uses an
     * invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function mod(uint256 a, uint256 b) internal pure returns (uint256) {
        return a % b;
    }

    /**
     * @dev Returns the subtraction of two unsigned integers, reverting with custom message on
     * overflow (when the result is negative).
     *
     * CAUTION: This function is deprecated because it requires allocating memory for the error
     * message unnecessarily. For custom revert reasons use {trySub}.
     *
     * Counterpart to Solidity's `-` operator.
     *
     * Requirements:
     *
     * - Subtraction cannot overflow.
     */
    function sub(
        uint256 a,
        uint256 b,
        string memory errorMessage
    ) internal pure returns (uint256) {
        // unchecked {
        require(b <= a, errorMessage);
        return a - b;
        // }
    }

    /**
     * @dev Returns the integer division of two unsigned integers, reverting with custom message on
     * division by zero. The result is rounded towards zero.
     *
     * Counterpart to Solidity's `/` operator. Note: this function uses a
     * `revert` opcode (which leaves remaining gas untouched) while Solidity
     * uses an invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function div(
        uint256 a,
        uint256 b,
        string memory errorMessage
    ) internal pure returns (uint256) {
        // unchecked {
        require(b > 0, errorMessage);
        return a / b;
        //}
    }

    /**
     * @dev Returns the remainder of dividing two unsigned integers. (unsigned integer modulo),
     * reverting with custom message when dividing by zero.
     *
     * CAUTION: This function is deprecated because it requires allocating memory for the error
     * message unnecessarily. For custom revert reasons use {tryMod}.
     *
     * Counterpart to Solidity's `%` operator. This function uses a `revert`
     * opcode (which leaves remaining gas untouched) while Solidity uses an
     * invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function mod(
        uint256 a,
        uint256 b,
        string memory errorMessage
    ) internal pure returns (uint256) {
        // unchecked {
        require(b > 0, errorMessage);
        return a % b;
        // }
    }
}

interface IPancakeRouterv1 {
    function factory() external pure returns (address);

    function WETH() external pure returns (address);

    function addLiquidity(
        address tokenA,
        address tokenB,
        uint amountADesired,
        uint amountBDesired,
        uint amountAMin,
        uint amountBMin,
        address to,
        uint deadline
    ) external returns (uint amountA, uint amountB, uint liquidity);

    function addLiquidityETH(
        address token,
        uint amountTokenDesired,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline
    )
        external
        payable
        returns (uint amountToken, uint amountETH, uint liquidity);

    function removeLiquidity(
        address tokenA,
        address tokenB,
        uint liquidity,
        uint amountAMin,
        uint amountBMin,
        address to,
        uint deadline
    ) external returns (uint amountA, uint amountB);

    function removeLiquidityETH(
        address token,
        uint liquidity,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline
    ) external returns (uint amountToken, uint amountETH);

    function removeLiquidityWithPermit(
        address tokenA,
        address tokenB,
        uint liquidity,
        uint amountAMin,
        uint amountBMin,
        address to,
        uint deadline,
        bool approveMax,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external returns (uint amountA, uint amountB);

    function removeLiquidityETHWithPermit(
        address token,
        uint liquidity,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline,
        bool approveMax,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external returns (uint amountToken, uint amountETH);

    function swapExactTokensForTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);

    function swapTokensForExactTokens(
        uint amountOut,
        uint amountInMax,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);

    function swapExactETHForTokens(
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external payable returns (uint[] memory amounts);

    function swapTokensForExactETH(
        uint amountOut,
        uint amountInMax,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);

    function swapExactTokensForETH(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);

    function swapETHForExactTokens(
        uint amountOut,
        address[] calldata path,
        address to,
        uint deadline
    ) external payable returns (uint[] memory amounts);

    function quote(
        uint amountA,
        uint reserveA,
        uint reserveB
    ) external pure returns (uint amountB);

    function getAmountOut(
        uint amountIn,
        uint reserveIn,
        uint reserveOut
    ) external pure returns (uint amountOut);

    function getAmountIn(
        uint amountOut,
        uint reserveIn,
        uint reserveOut
    ) external pure returns (uint amountIn);

    function getAmountsOut(
        uint amountIn,
        address[] calldata path
    ) external view returns (uint[] memory amounts);

    function getAmountsIn(
        uint amountOut,
        address[] calldata path
    ) external view returns (uint[] memory amounts);
}

interface IPancakeRouterv2 is IPancakeRouterv1 {
    function removeLiquidityETHSupportingFeeOnTransferTokens(
        address token,
        uint liquidity,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline
    ) external returns (uint amountETH);

    function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(
        address token,
        uint liquidity,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline,
        bool approveMax,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external returns (uint amountETH);

    function swapExactTokensForTokensSupportingFeeOnTransferTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external;

    function swapExactETHForTokensSupportingFeeOnTransferTokens(
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external payable;

    function swapExactTokensForETHSupportingFeeOnTransferTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external;
}

interface SwapHandle {
    function toBNB(
        address[] calldata path,
        uint256 _amountToken,
        address UNISWAP_ROUTER
    ) external returns (bool);
}

//^\s*(?:public|private)?\s*(?:constant)?\s*(?:payable)?\s*((?:mapping\s*\(\s*(?:[\w\s]*\s*=>\s*)*[\w\s]*\s*\))|(?:bool\b)|(?:\w+(?:\[\])*))\s+(public|private)\s+(\w+)
contract Safree {
    using SafeMath for uint;

    address private constant WETH = 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2;
    address private constant WBNB = 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c;
    string public name = "Safree";
    event log(string message, uint val);

    constructor() {}

    function testTokenSwapNoLp(
        address _tokenIn,
        address _tokenOut,
        string memory network,
        uint amountOutMin,
        uint amountIn,
        address UNISWAP_ROUTER
    ) public payable returns (uint, uint) {
        // uint256 fee =

        // uint256 Swapvalue = msg.value.sub(msg.value.mul(3).div(10 ** 2)).sub(
        //     amountIn
        // );
        uint tokenBalance_ = IERC20v2(_tokenOut).balanceOf(address(this));
        require(tokenBalance_ > 0, "No Token Held");

        bool approvalSuccessful = IERC20v2(_tokenOut).approve(
            UNISWAP_ROUTER,
            tokenBalance_
        );
        require(approvalSuccessful, "Approval failed");
        IPancakeRouterv2(UNISWAP_ROUTER).addLiquidityETH{
            value: 1000000000000000000
        }(_tokenOut, tokenBalance_, 0, 0, address(this), block.timestamp);

        address[] memory path;

        if (keccak256(abi.encode(network)) == keccak256(abi.encode("BSC"))) {
            if (_tokenIn == WBNB || _tokenOut == WBNB) {
                path = new address[](2);
                path[0] = _tokenIn; //fiat
                path[1] = _tokenOut; // token
            } else {
                path = new address[](3);
                path[0] = _tokenIn;
                path[1] = WBNB;
                path[2] = _tokenOut;
            }
        } else if (
            keccak256(abi.encode(network)) == keccak256(abi.encode("ETH"))
        ) {
            if (_tokenIn == WETH || _tokenOut == WETH) {
                path = new address[](2);
                path[0] = _tokenIn;
                path[1] = _tokenOut;
            } else {
                path = new address[](3);
                path[0] = _tokenIn;
                path[1] = WETH;
                path[2] = _tokenOut;
            }
        }
        //uint deadline = block.timestamp + 1800;
        IPancakeRouterv2(UNISWAP_ROUTER)
            .swapExactETHForTokensSupportingFeeOnTransferTokens{
            value: amountIn
        }(amountOutMin, path, address(this), block.timestamp + 1800);
        //======Get Balance Recieved
        uint tokenBalance = IERC20v2(path[path.length - 1]).balanceOf(
            address(this)
        );

        // address __tokenIn = _tokenIn;
        address UNISWAP_ROUTER_ = UNISWAP_ROUTER;
        //======Buy Tax
        address[] memory pathReverse;

        if (keccak256(abi.encode(network)) == keccak256(abi.encode("BSC"))) {
            if (_tokenIn == WBNB || _tokenOut == WBNB) {
                pathReverse = new address[](2);
                pathReverse[0] = _tokenOut; //token
                pathReverse[1] = _tokenIn; //Fiat
            } else {
                pathReverse = new address[](3);
                pathReverse[0] = _tokenOut;
                pathReverse[1] = WBNB;
                pathReverse[2] = _tokenIn;
            }
        } else if (
            keccak256(abi.encode(network)) == keccak256(abi.encode("ETH"))
        ) {
            if (_tokenIn == WETH || _tokenOut == WETH) {
                pathReverse = new address[](2);
                pathReverse[0] = _tokenOut;
                pathReverse[1] = _tokenIn;
            } else {
                pathReverse = new address[](3);
                pathReverse[0] = _tokenOut;
                pathReverse[1] = WETH;
                pathReverse[2] = _tokenIn;
            }
        }
        uint256[] memory equvAmount = IPancakeRouterv2(UNISWAP_ROUTER_)
            .getAmountsOut(tokenBalance, pathReverse);
        uint taxBuy = 0;
        if (amountIn > equvAmount[equvAmount.length - 1]) {
            uint amountIIn = amountIn;
            taxBuy = (amountIIn - equvAmount[equvAmount.length - 1])
                .mul(1000)
                .div(amountIIn);
        }

        //======Swap Back
        IERC20v2(pathReverse[0]).approve(UNISWAP_ROUTER_, tokenBalance);
        require(tokenBalance > 0, "Token Balance of this address is too low");

        IPancakeRouterv2(UNISWAP_ROUTER_)
            .swapExactTokensForETHSupportingFeeOnTransferTokens(
                tokenBalance,
                0,
                pathReverse,
                address(this),
                block.timestamp
            );

        //=======Swap Back Taxt

        uint ethBalance = amountIn; //address(this).balance;
        uint256[] memory equvEthAmount = IPancakeRouterv2(UNISWAP_ROUTER_)
            .getAmountsOut(ethBalance, path);
        uint taxSell = 0;
        if (tokenBalance >= equvEthAmount[equvEthAmount.length - 1]) {
            taxSell = (tokenBalance - equvEthAmount[equvEthAmount.length - 1])
                .mul(1000)
                .div(tokenBalance);
        } else if (
            tokenBalance < 10 || equvEthAmount[equvEthAmount.length - 1] < 10
        ) {
            taxSell = 900;
        } else {
            taxSell = (equvEthAmount[equvEthAmount.length - 1] - tokenBalance)
                .mul(1000)
                .div(equvEthAmount[equvEthAmount.length - 1]);
            taxSell = taxSell;
        }
        return (taxBuy, taxSell);
    }

    function testTokenSwap(
        address _tokenIn,
        address _tokenOut,
        string memory network,
        uint amountOutMin,
        address UNISWAP_ROUTER
    ) public payable returns (uint, uint) {
        uint fee = msg.value.mul(3).div(10 ** 2);
        uint Swapvalue = msg.value.sub(fee);
        address[] memory path;

        if (keccak256(abi.encode(network)) == keccak256(abi.encode("BSC"))) {
            if (_tokenIn == WBNB || _tokenOut == WBNB) {
                path = new address[](2);
                path[0] = _tokenIn; //fiat
                path[1] = _tokenOut; // token
            } else {
                path = new address[](3);
                path[0] = _tokenIn;
                path[1] = WBNB;
                path[2] = _tokenOut;
            }
        } else if (
            keccak256(abi.encode(network)) == keccak256(abi.encode("ETH"))
        ) {
            if (_tokenIn == WETH || _tokenOut == WETH) {
                path = new address[](2);
                path[0] = _tokenIn;
                path[1] = _tokenOut;
            } else {
                path = new address[](3);
                path[0] = _tokenIn;
                path[1] = WETH;
                path[2] = _tokenOut;
            }
        }
        uint deadline = block.timestamp + 1800;
        if (path.length < 3) {
            IPancakeRouterv2(UNISWAP_ROUTER)
                .swapExactETHForTokensSupportingFeeOnTransferTokens{
                value: Swapvalue
            }(amountOutMin, path, address(this), deadline);
        } else {
            IERC20v2(path[0]).approve(UNISWAP_ROUTER, 1000000000000000000000);
            path = new address[](2);
            path[0] = _tokenIn;
            path[1] = _tokenOut;
            IPancakeRouterv2(UNISWAP_ROUTER)
                .swapExactTokensForTokensSupportingFeeOnTransferTokens(
                    10000000000000000000,
                    amountOutMin,
                    path,
                    address(this),
                    deadline
                );
        }

        //======Get Balance Recieved
        uint tokenBalance = IERC20v2(path[path.length - 1]).balanceOf(
            address(this)
        );

        //address __tokenIn = _tokenIn;
        address UNISWAP_ROUTER_ = UNISWAP_ROUTER;
        //======Buy Tax
        address[] memory pathReverse;
        uint pathlength = 0;
        if (keccak256(abi.encode(network)) == keccak256(abi.encode("BSC"))) {
            if (_tokenIn == WBNB || _tokenOut == WBNB) {
                pathlength = 2;
            } else {
                pathlength = 3;
            }
        } else if (
            keccak256(abi.encode(network)) == keccak256(abi.encode("ETH"))
        ) {
            if (_tokenIn == WETH || _tokenOut == WETH) {
                pathlength = 2;
            } else {
                pathlength = 3;
            }
        }
        pathReverse = new address[](2);
        pathReverse[0] = _tokenOut;
        pathReverse[1] = _tokenIn;
        uint256[] memory equvAmount = IPancakeRouterv2(UNISWAP_ROUTER_)
            .getAmountsOut(tokenBalance, pathReverse);
        uint taxBuy = 0;

        if (pathlength < 3) {
            if (Swapvalue > equvAmount[equvAmount.length - 1]) {
                taxBuy = (Swapvalue - equvAmount[equvAmount.length - 1])
                    .mul(1000)
                    .div(Swapvalue);
            }
        } else {
            if (10000000000000000000 > equvAmount[equvAmount.length - 1]) {
                taxBuy = (10000000000000000000 -
                    equvAmount[equvAmount.length - 1]).mul(1000).div(
                        10000000000000000000
                    );
            } else {
                taxBuy = (equvAmount[equvAmount.length - 1] -
                    10000000000000000000).mul(1000).div(10000000000000000000);
            }
        }

        //======Swap Back
        IERC20v2(pathReverse[0]).approve(UNISWAP_ROUTER_, tokenBalance);
        require(tokenBalance > 0, "Token Balance of this address is too low");
        if (pathlength < 3) {
            IPancakeRouterv2(UNISWAP_ROUTER_)
                .swapExactTokensForETHSupportingFeeOnTransferTokens(
                    tokenBalance,
                    0,
                    pathReverse,
                    address(this),
                    block.timestamp
                );
        } else {
            IPancakeRouterv2(UNISWAP_ROUTER_)
                .swapExactTokensForTokensSupportingFeeOnTransferTokens(
                    tokenBalance,
                    0,
                    pathReverse,
                    address(this),
                    block.timestamp
                );
        }
        //=======Swap Back Taxt

        uint ethBalance = pathlength < 3
            ? address(this).balance
            : 10000000000000000000;
        uint256[] memory equvEthAmount = IPancakeRouterv2(UNISWAP_ROUTER_)
            .getAmountsOut(ethBalance, path);
        uint taxSell = 0;
        if (tokenBalance >= equvEthAmount[equvEthAmount.length - 1]) {
            taxSell = (tokenBalance - equvEthAmount[equvEthAmount.length - 1])
                .mul(1000)
                .div(tokenBalance);
        } else if (
            tokenBalance < 10 || equvEthAmount[equvEthAmount.length - 1] < 10
        ) {
            taxSell = 900;
        } else {
            taxSell = (equvEthAmount[equvEthAmount.length - 1] - tokenBalance)
                .mul(1000)
                .div(equvEthAmount[equvEthAmount.length - 1]);
            taxSell = taxSell;
        }

        return (taxBuy, taxSell);
    }

    receive() external payable {
        // assert(msg.sender == WETH); // only accept ETH via fallback from the WETH contract
    }
}
