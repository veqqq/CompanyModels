using Pkg
Pkg.add("LibPQ")
Pkg.add("Plots")
Pkg.add("DataFrames")


using LibPQ
using Plots
using DataFrames

function fetch_commodity_data(conn::LibPQ.Connection)
    query = """
    SELECT
    d.date,
    d.open,
    d.TickerSymbol
FROM
    dailyohlcvs d
WHERE
    d.TickerSymbol IN (
        SELECT
            t.TickerSymbol
        FROM
            tickers t
        WHERE
            t.type LIKE '%coal%')
    """


    
    result = LibPQ.execute(conn, query)
    
    # Print the result obtained from the query
    println("Result from query: ", result)
    
    dates = []
    values = []
    tickers = []
    
    for row in result
        push!(dates, row[1])
        push!(values, row[2])
        push!(tickers, row[3])
    end
    
    return dates, values, tickers
end
function plot_commodity_data(dates, values, tickers)
    df = DataFrame("Date" => dates, "Value" => values, "Ticker" => tickers)
    grouped_data = groupby(df, :Ticker)
    
    p = plot(size=(2000, 1200))

    # Use a colormap for different colors for each ticker
    colors = Plots.get_color_palette(:auto, length(grouped_data))

    for (i, group) in enumerate(grouped_data)
        dates = group[!, "Date"]
        values = group[!, "Value"]
        color = colors[i]
        plot!(p, dates, values, line=:auto, color=color, seriestype=:line, label=string(group[1, "Ticker"]))
    end
    
    plot!(p, xlabel="Date", ylabel="Stock Price")
    display(p)
end




function main()
    # Replace with your PostgreSQL connection parameters
    conn = LibPQ.Connection("host=127.0.0.1 dbname=financial_markets user=postgres password=password2")
    
    try
        dates, values, tickers = fetch_commodity_data(conn)
        plot_commodity_data(dates, values, tickers)
    finally
        LibPQ.close(conn)
    end
end

main()