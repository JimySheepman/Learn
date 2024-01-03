#include <array>
#include <iostream>
#include <string_view>
#include <tuple>
#include <type_traits>

namespace a::b::c
{
    inline constexpr std::string_view str{ "hello"};
} // namespace a::b::c

template <typename... T>
std::tuple<std::size_t, std::common_type<T...>> sum(T ...args)
{
return { sizeof...(T), (args+...)};
}

int main()
{
    auto[iNumbers, iSum]{sum(1,2,3)};
    std::cout << a::b::c::str << ' '<< iNumbers << ' ' << iSum << '\n';
    std::array arr{1,2,3};

std::cout << std::size(arr)<< '\n';

return 0;
}
