use std::error::Error;

pub async fn scrape(url: &str, selector: String) -> Result<Vec<String>, Box<dyn Error>> {
    let html = reqwest::get(url).await?.text().await?;
    let document = scraper::Html::parse_document(&html);
    let title_selector = scraper::Selector::parse(&selector).unwrap();
    let titles = document.select(&title_selector).map(|x| x.inner_html());

    let data = titles.map(|item| item).collect::<Vec<String>>();
    Ok(data)
}
