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
    
    p = plot()

    # Choose a line style, color, and symbol for all plots
    line_style = :solid
    line_color = :blue
    line_symbol = :circle
    
    for group in grouped_data
        ticker = group[!, 1]
        dates = group[!, "Date"]  # Accessing the "Date" column directly
        values = group[!, "Value"]  # Accessing the "Value" column directly
        plot!(p, dates, values, label=ticker, line=:auto, color=line_color, seriestype=:scatter, marker=:auto, linestyle=line_style)
    end
    
    plot!(p, xlabel="Date", ylabel="Commodity Value", legend=:topright)
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